import { Box, VStack } from "@chakra-ui/layout"
import { Heading, Button } from "@chakra-ui/react"
import { FormControl, FormLabel, Input, Select } from "@chakra-ui/react"
import { useState } from "react";
import CreatableSelect from 'react-select/creatable';
import addAgentMutation, {AgentType} from '../Services/AddAgent'

interface HomeProps {

}

const AddAgent:React.FC<HomeProps> = (props) => {
  const [state, setState] = useState<AgentType>({name: "", ult: "", class: ""})
  const mutation = addAgentMutation()

  const submit = () => {
    mutation.mutate(state)
  }
  
  return (
    <Box p={10}>
      <VStack spacing={8}>
        <Heading size='xl'>Add Agent</Heading>
        <FormControl>
          <FormLabel>Name</FormLabel>
          <Input variant="outline" value={state.name} onChange={(e) => setState({...state, name: e.target.value})} />
        </FormControl>
        <FormControl>
          <FormLabel>Ult</FormLabel>
          <Input variant="outline" value={state.ult} onChange={(e) => setState({...state, ult: e.target.value})} />
        </FormControl>
        <FormControl>
          <FormLabel>Class</FormLabel>
          <Input variant="outline" value={state.class} onChange={(e) => setState({...state, class: e.target.value})} />
        </FormControl>
        <Button colorScheme='brand' onClick={submit}>Add</Button>
      </VStack>
    </Box>
  )
}

export default AddAgent