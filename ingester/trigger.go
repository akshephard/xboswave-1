package main

import (
	"github.com/gtfierro/xboswave/ingester/types"
	logrus "github.com/sirupsen/logrus"
)

type subscriptionChange struct {
	add bool
	uri types.SubscriptionURI
}

func (ingest *Ingester) handleSubscriptionChanges() {
	go func() {
		for subChg := range ingest.pendingSubs {
			logrus.Infof("Processing subscription change: adding? %v uri: %v", subChg.add, subChg.uri)
			ingest.subsLock.Lock()
			sub, found := ingest.subs[subChg.uri]
			if subChg.add && !found {
				sub, err := ingest.newSubscription(subChg.uri)
				if err != nil {
					logrus.Error("Error processing subscription: %v", err)
					//register this error with the subscription so we can go find it later
					if serr := ingest.cfgmgr.MarkErrorURI(subChg.uri, err.Error()); serr != nil {
						logrus.Warning(serr)
					}
					// through the CLI
				} else {
					if serr := ingest.cfgmgr.ClearErrorURI(subChg.uri); serr != nil {
						logrus.Warning(serr)
					}
					ingest.subs[subChg.uri] = sub
					sub.timer.Reset()
				}
				// handle add subscription
			} else if !subChg.add && found {
				// handle delete
				sub.stop <- struct{}{}
				delete(ingest.subs, subChg.uri)
			}

			ingest.subsLock.Unlock()
		}
	}()
}
