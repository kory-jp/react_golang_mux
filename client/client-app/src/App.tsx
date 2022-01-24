import { Provider } from 'react-redux';
import InputArea from './components/organisms/posts/InputArea';
import createInitStore from './reducks/store/store';
import Router from './router/Router';

const store = createInitStore()

function App() {
  return (
    <>
      <Provider store={store}>
        <Router />
      </Provider>
    </>
  );
}

export default App;
