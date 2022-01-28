
import Axios from 'axios'
import { GraphQLClient, gql } from 'graphql-request'

const BASE_URL = `http://${process.env.REACT_APP_BASE_URL}:${process.env.REACT_APP_API_PORT}`

const axios = Axios.create({
  baseURL: `${process.env.REACT_APP_BASE_URL}:${process.env.REACT_APP_API_PORT}`
})

const graphql = new GraphQLClient(`${BASE_URL}/graphql`)

export {
  axios,
  graphql
}