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
        console.log("Listing Logging");
        const  {listAgents: data}  = await graphql.request(query)
        console.log("Listing ", data)
        return data
    }, {
        refetchOnWindowFocus: false
    })
}

export default useListAgents