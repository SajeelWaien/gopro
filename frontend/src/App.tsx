import './App.css';
import { ReactLocation, Router } from 'react-location'
import { Container } from '@chakra-ui/layout';
import Home from './pages/Home'
import Topbar from './components/Topbar';
import AddAgent from './pages/AddAgent';

const location = new ReactLocation()

function App() {
  return (
    <>
      <Topbar />
      <Container pt={10}>
        <Router
          location={location}
          routes={[
            { path: "/", element: <Home /> },
            { path: "/add-agent", element: <AddAgent /> }
          ]}
        />
      </Container>
    </>
  );
}

export default App;
