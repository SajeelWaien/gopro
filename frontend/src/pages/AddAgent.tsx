import { Box, VStack } from "@chakra-ui/layout"
import { Heading } from "@chakra-ui/react"
import { FormControl, FormLabel, Input, Select } from "@chakra-ui/react"
import CreatableSelect from 'react-select/creatable';

interface HomeProps {

}

const AddAgent:React.FC<HomeProps> = (props) => {
  return (
    <Box p={10}>
      <VStack spacing={8}>
        <Heading size='xl'>Add Agent</Heading>
        <FormControl>
          <FormLabel>Name</FormLabel>
          <Input variant="outline" />
        </FormControl>
        <FormControl>
          <FormLabel>Ult</FormLabel>
          <Input  />
        </FormControl>
        <FormControl>
          <FormLabel>Abilities</FormLabel>
          <Select variant="filled" />
        </FormControl>
      </VStack>
    </Box>
  )
}

export default AddAgent