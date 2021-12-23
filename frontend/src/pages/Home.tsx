import { Box, VStack } from "@chakra-ui/layout"
import { Heading } from "@chakra-ui/react"
import { FormControl, FormLabel, Input, Select, Center } from "@chakra-ui/react"
import { Link } from "react-location";
import CreatableSelect from 'react-select/creatable';

interface HomeProps {

}

const Home:React.FC<HomeProps> = (props) => {
  return (
    <Box p={10}>
      <Center>
        <Heading size='xl'>Home</Heading>
        <Link to="add-agent">Home</Link>
      </Center>
    </Box>
  )
}

export default Home