import { gql } from 'graphql-request'
import { useMutation } from 'react-query'
import { graphql, axios } from '../config/request'
// import axios from 'axios'

export type AgentType = {
  name: string,
  ult: string,
  class: string
}

const AddAgentMutation = (/* variables: AgentType */) => {
  const mutation = gql`
    mutation AddAgent($name: String!, $ult: String!, $class: Class!){
      addAgent(object: {name: $name, ult: $ult, class: $class}) {
        name
        ult
        class
      }
    }
  `

  return useMutation((newAgent: AgentType) => {
    console.log(newAgent)
    return graphql.request(mutation, newAgent)
  })
}

export default AddAgentMutation