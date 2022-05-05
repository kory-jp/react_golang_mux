import 'normalize.css'
import { ConnectedRouter } from 'connected-react-router';
import { Provider } from 'react-redux';
import createInitStore from './reducks/store/store';
import Router from './router/Router';
import * as History from "history";
import { IncompleteTaskCardCountProvider } from './providers/IncompleteTaskCardCount';

const history = History.createBrowserHistory();
const store = createInitStore(history);

function App() {
  return (
    <>
      <Provider store={store}>
        <IncompleteTaskCardCountProvider>
          <ConnectedRouter history={history}>
            <Router />
          </ConnectedRouter>
        </IncompleteTaskCardCountProvider>
      </Provider>
    </>
  );
}

export default App;
