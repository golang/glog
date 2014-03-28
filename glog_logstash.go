package glog

import (
    "encoding/json"
    "fmt"
    "net"
    "os"
    "strings"
    "time"
)

type logstashMessage struct {
    Type    string `json:"type"`
    Message string `json:"message"`
}

// manageLogstashConnection manages the connection to the logstash server.
func (l *loggingT) manageLogstashConnection() {
    var err error
    for {
        select {
        case _ = <-l.logstashStop:
            return
        default:
            if l.logstashConn == nil {
                fmt.Fprintln(os.Stderr, "Trying to connect to logstash server...")
                l.logstashConn, err = net.Dial("tcp", l.logstashURL)
                if err != nil {
                    l.logstashConn = nil
                } else {
                    fmt.Fprintln(os.Stderr, "Connected to logstash server.")
                }
            }
            time.Sleep(time.Second)
        }
    }
}

// handleLogstashMessages sends logs to logstash.
func (l *loggingT) handleLogstashMessages() {
    for {
        select {
        case _ = <-l.logstashStop:
            return
        case data := <-l.logstashChan:
            lm := logstashMessage{}
            lm.Type = l.logstashType
            lm.Message = strings.TrimSpace(data)
            packet, err := json.Marshal(lm)
            if err != nil {
                fmt.Fprintln(os.Stderr, "Failed to marshal logstashMessage.")
                continue
            } else {
                if l.logstashConn != nil {
                    _, err := fmt.Fprintln(l.logstashConn, string(packet))
                    if err != nil {
                        fmt.Fprintln(os.Stderr, "Not connected to logstash server, attempting reconnect.")
                        l.logstashConn = nil
                        continue
                    }
                } else {
                    // There is no connection, so the log line is dropped.
                    // Might be nice to add a buffer here so that we can ship
                    // logs after the connection is up.
                }
            }
        }
    }
}

// startLogstash creates the logstash channel and kicks off the connection and message handlers.
func (l *loggingT) startLogstash() {
    l.logstashChan = make(chan string)
    go l.manageLogstashConnection()
    go l.handleLogstashMessages()
}

// stopLogstash signals the goroutines manageLogstashConnection and handleLogstashMessages to exit.
func (l *loggingT) stopLogstash() {
    l.logstashStop <- true
    l.logstashStop <- true
}

