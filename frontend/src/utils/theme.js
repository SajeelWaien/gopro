import { extendTheme } from '@chakra-ui/react'

const PRIMARY_COLOR = '#ff4655'

const inputVariantOutlined = () => ({
  field: {
    _focus: {
      borderColor: "var(--chakra-ui-focus-ring-color)",
      boxShadow: "0 0 0 2px var(--chakra-ui-focus-ring-color)"
    }
  }
})

const theme = extendTheme({
  colors: {
    brand: PRIMARY_COLOR,
  },
  styles: {
    global: {
      ":host, :root": {
        "--chakra-ui-focus-ring-color": PRIMARY_COLOR
      },
      'html, body': {
        backgroundColor: '#081229',
        color: 'white'
      }
    },
  },
  fonts: {
    heading: "'Spectral SC', serif",
    body: "'Spectral SC', serif",
  },
  components: {
    Input: {
      variants: {
        outline: inputVariantOutlined
      }
    },
    Heading: {
      baseStyle: {
        textDecoration: 'underline',
        textDecorationColor: 'brand',
        textUnderlineOffset: '10px'
      },
    }
  }
})

export default theme