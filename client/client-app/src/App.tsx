import { Provider } from 'react-redux';
import InputArea from './components/organisms/posts/InputArea';
import createInitStore from './reducks/store/store';

const store = createInitStore()

function App() {
  return (
    <>
      <Provider store={store}>
        <InputArea />
      </Provider>
    </>
  );
}

export default App;
