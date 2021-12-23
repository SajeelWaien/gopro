import { Box, Center, HStack } from "@chakra-ui/react"
import { Link } from "react-location"

const Topbar = () => {
  return (
    <Box>
      <Box width='100%' height='50px' bgColor='brand' position='absolute' opacity='0.7'>
        <Center>
          <HStack>
            <Box>
              {/* <Link
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
              </Link> */}
            </Box>
          </HStack>
        </Center>
      </Box>
      <Box position='absolute' sx={{ filter: 'drop-shadow(4px -2px 1px rgba(0,0,0,0.5))'}}>
        <Box
          width='80px'
          height={200}
          bgColor='brand'
          sx={{ clipPath: "polygon(0px 0px, 100% 50%, 100% 100%, 0 50%)" }}
        >

        </Box>
      </Box>
      <Box position='absolute' right={0} sx={{ filter: 'drop-shadow(-4px -2px 1px rgba(0,0,0,0.5))'}}>
        <Box
          width='80px'
          height={200}
          bgColor='brand'
          transform={'scaleY(-1)'}
          sx={{ clipPath: "polygon(0px 0px, 100% 50%, 100% 100%, 0 50%)" }}
        >

        </Box>
      </Box>
    </Box>
  )
}

export default Topbar