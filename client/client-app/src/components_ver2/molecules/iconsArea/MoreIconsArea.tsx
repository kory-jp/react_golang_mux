import { VFC } from "react";
import DefaultArea from "./DefaultArea";
import StickyNote2Icon from '@mui/icons-material/StickyNote2';

type Props = {
  finish: boolean,
  onChangeIsFinished: () => void,
  onClickDelete: () => void,
  onClickToShowTodo: () => void,
}

export const MoreIconsArea: VFC<Props> = (props) => {
  const {finish, onChangeIsFinished, onClickDelete, onClickToShowTodo} = props

  return(
    <>
      <DefaultArea
        finish={finish}
        onChangeIsFinished={onChangeIsFinished}
        onClickDelete={onClickDelete}
        onClickSomeEvent={onClickToShowTodo}
      >
        <StickyNote2Icon 
          sx={{
            color: '#587FBA',
          }}
        />
        more       
      </DefaultArea>
    </>
  )
}

export default MoreIconsArea;