package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/models"
)

var classesEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Class",
	Description: "Class to which a character belongs to based on abilities and playing style",
	Values: graphql.EnumValueConfigMap{
		"Sentinel": &graphql.EnumValueConfig{
			Value: 1,
		},
		"Duelist": &graphql.EnumValueConfig{
			Value: 2,
		},
		"Initiator": &graphql.EnumValueConfig{
			Value: 3,
		},
		"Controller": &graphql.EnumValueConfig{
			Value: 4,
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
			Type: graphql.NewList(classesEnum),
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
					Type: graphql.NewList(classesEnum),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				ult, _ := params.Args["ult"].(string)
				class, _ := params.Args["class"].(string)

				agent := models.Agent{Name: name, Ult: ult, Class: class}

				result := database.DBCon.Create(&agent)

				return result, nil
			},
		},
	},
})

var AgentSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Mutation: rootMutation,
})
