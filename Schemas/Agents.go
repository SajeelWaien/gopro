package schemas

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/models"
)

var classesEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Class",
	Description: "Class to which a character belongs to based on abilities and playing style",
	Values: graphql.EnumValueConfigMap{
		"Sentinel": &graphql.EnumValueConfig{
			Value: "Sentinel",
		},
		"Duelist": &graphql.EnumValueConfig{
			Value: "Duelist",
		},
		"Initiator": &graphql.EnumValueConfig{
			Value: "Initiator",
		},
		"Controller": &graphql.EnumValueConfig{
			Value: "Controller",
		},
	},
})

var agentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Agent",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"ult": &graphql.Field{
			Type: graphql.String,
		},
		"class": &graphql.Field{
			Type: classesEnum,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"listAgents": &graphql.Field{
			Type:        graphql.NewList(agentType),
			Description: "Returns a list of all agents.",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var agents []models.Agent
				result := database.DBCon.Find(&agents)

				if result.Error != nil {
					return nil, result.Error
				}

				return &agents, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addAgent": &graphql.Field{
			Type:        agentType,
			Description: "Add a new agent",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"ult": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"class": &graphql.ArgumentConfig{
					Type: classesEnum,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				ult, _ := params.Args["ult"].(string)
				class, _ := params.Args["class"].(string)

				fmt.Printf("{}, {}, {}", name, ult, class)

				agent := models.Agent{Name: name, Ult: ult, Class: class}

				result := database.DBCon.Create(&agent)

				if result.Error != nil {
					return nil, result.Error
				}

				return &agent, nil
			},
		},
	},
})

var AgentSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
