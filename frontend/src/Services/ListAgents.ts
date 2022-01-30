import { gql } from "graphql-request"
import { useQuery } from "react-query"
import { graphql } from "../config/request"

function useListAgents() {
    return useQuery('listAgents', async () => {
        const query = gql`
            query ListAgents {
                listAgents {
                    name
                    class
                    ult
                }
            }
        `
        const { listAgents } = await graphql.request(query)
        console.log("@@@@@ ", listAgents)
        return listAgents
    })
}

export default useListAgents