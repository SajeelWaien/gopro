import { gql } from 'graphql-request'
import { useMutation } from 'react-query'
import { graphql } from '../config/request'

export type AgentType = {
  name: string,
  ult: string,
  class: string[]
}

const AddAgentMutation = (/* variables: AgentType */) => {
  const mutation = gql`
    mutation AddAgent($name: String!, $ult: String!, $class: String!){
      addAgent(name: $name, ult: $ult, class: $class) {
        name
        ult
        class
      }
    }
  `

  return useMutation((newAgent: AgentType) => {
    return graphql.request(mutation, newAgent)
  })
}

export default AddAgentMutation