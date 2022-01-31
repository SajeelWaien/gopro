import { Box, Center, HStack } from "@chakra-ui/react"
import { Link } from "react-location"

const Topbar = () => {
  return (
    <Box position='absolute' width='100%' left={0} top={0}>
      <Box width='100%' height='50px' bgColor='brand.500' position='absolute' opacity='0.7'>
        <Center height={'100%'}>
          <HStack spacing={5}>
            <Box>
              <Link
                to={'add-agent'}
                className={``}
                // // Make "active" links bold
                // getActiveProps={() => ({ className: tw`font-bold` })}
                // activeOptions={{
                //   // If the route points to the root of it's parent,
                //   // make sure it's only active if it's exact
                //   exact: to === ".",
                // }}
              >
                Add New
              </Link>
            </Box>
            <Box>
              <Link to={'list-agents'} className={``}>
                List Agents
              </Link>
            </Box>
          </HStack>
        </Center>
      </Box>
      <Box position='absolute' sx={{ filter: 'drop-shadow(4px -2px 1px rgba(0,0,0,0.5))'}}>
        <Box
          width='80px'
          height={200}
          bgColor='brand.500'
          sx={{ clipPath: "polygon(0px 0px, 100% 50%, 100% 100%, 0 50%)" }}
        >

        </Box>
      </Box>
      <Box position='absolute' right={0} sx={{ filter: 'drop-shadow(-4px -2px 1px rgba(0,0,0,0.5))'}}>
        <Box
          width='80px'
          height={200}
          bgColor='brand.500'
          transform={'scaleY(-1)'}
          sx={{ clipPath: "polygon(0px 0px, 100% 50%, 100% 100%, 0 50%)" }}
        >

        </Box>
      </Box>
    </Box>
  )
}

export default Topbar