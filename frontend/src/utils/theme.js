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
    brand: {
      100: '#FF9AA2',
      200: '#FF8993',
      300: '#FF7884',
      400: '#FF6874',
      500: PRIMARY_COLOR,
      600: '#E64E5B',
      700: '#CC4551',
      800: '#B33D47',
      900: '#99343C',
    }
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
        textDecorationColor: 'brand.500',
        textUnderlineOffset: '10px'
      },
    }
  }
})

export default theme