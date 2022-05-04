import 'normalize.css'
import { ConnectedRouter } from 'connected-react-router';
import { Provider } from 'react-redux';
import createInitStore from './reducks/store/store';
import Router from './router/Router';
import * as History from "history";

const history = History.createBrowserHistory();
const store = createInitStore(history);

function App() {
  return (
    <>
      <Provider store={store}>
         <ConnectedRouter history={history}>
           <Router />
         </ConnectedRouter>
      </Provider>
    </>
  );
}

export default App;
