import './App.css';
import { Outlet, ReactLocation, Router } from 'react-location'
import { Container } from '@chakra-ui/layout';
import Home from './pages/Home'
import Topbar from './components/Topbar';
import AddAgent from './pages/AddAgent';
import ListAgents from './pages/ListAgents';

const location = new ReactLocation()

function App() {
  return (
    <>
      <Container pt={20}>
        <Router
          location={location}
          routes={[
            { path: "/", element: <Home /> },
            { path: "/add-agent", element: <AddAgent /> },
            { path: "/list-agents", element: <ListAgents /> }
          ]}
        >
          <Topbar />
          <Outlet />
        </Router>
      </Container>
    </>
  );
}

export default App;
