package indexer

import (

	"github.com/antholord/poeIndexer/api"
	"github.com/antholord/poeIndexer/subscription"
	"github.com/mailru/easyjson"
)

func processStash(stash *api.Stash, m *subscription.Manager) {
	//log.Println(len(m.SubMap))
	for _, item := range stash.Items {

		for itemSearch, clients := range m.SubMap {
			if (matchesCriterias(itemSearch, &item)){
				go broadcast(clients, api.ItemResult{item, stash.AccountName, stash.LastCharacterName, stash.Id, stash.Label, stash.Type, "",})
			}
		}

		/*
		if item.Type == "Ancient Reliquary Key" {
			//result :=
			//log.Printf(msg)
			output <- api.ItemResult{item, stash.AccountName, stash.LastCharacterName, stash.Id, stash.Label, stash.Type, "",}
		}*/
	}
}

func matchesCriterias(s *subscription.ItemSearch, item *api.Item) bool{

	var match bool = true

	if (s.League != "" && s.League == item.League){
		if (s.Type != "" && s.Type == item.Type){
			match = true
		}
		if (s.Name != "" && s.Name == item.Name){
			match = true
		}
		return match;
	}else{
		return false;
	}
}

func broadcast(clients map[*subscription.Client]bool, s api.ItemResult){
	for client,_:= range clients {
		json, _ := easyjson.Marshal(s)
		client.Send <- json
	}
}