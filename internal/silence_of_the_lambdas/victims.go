package silenceofthelambdas

import log "github.com/sirupsen/logrus"

type Victim struct {
	Killer       string `json:"killer"`
	FullName     string `json:"full_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Actor        string `json:"actor"`
	Movie        string `json:"movie"`
	CauseOfDeath string `json:"cause_of_death"`
	Occupation   string `json:"occupation"`
	Cannibalized bool   `json:"cannibalized"`
	Image        string `json:"image"`
}

type Victims struct {
	Victims []*Victim
}

func (c *SilenceOfTheLambdasClient) ReturnVictimsByKiller(killerName string, tableName string) (*Victims, error) {
	fields := log.Fields{"killer": killerName, "table_name": tableName}

	victims := &Victims{Victims: make([]*Victim, 0)}
	dynamoVictims, err := c.db.ReturnVictimsByKiller(killerName, tableName)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING VICTIMS FROM DYNAMODB: %+v", err)
		return nil, err
	}

	for _, victim := range dynamoVictims.Victims {
		newVictim := &Victim{
			Killer:       victim.Killer,
			FullName:     victim.FullName,
			FirstName:    victim.FirstName,
			LastName:     victim.LastName,
			Actor:        victim.Actor,
			Movie:        victim.Movie,
			CauseOfDeath: victim.CauseOfDeath,
			Occupation:   victim.Occupation,
			Cannibalized: victim.Cannibalized,
			Image:        victim.Image,
		}
		victims.Victims = append(victims.Victims, newVictim)
	}

	return victims, nil
}
