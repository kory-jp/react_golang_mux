import { VFC } from "react";
import DefaultArea from "./DefaultArea";
import EditIcon from '@mui/icons-material/Edit';

type Props = {
  finish: boolean,
  onChangeIsFinished: () => void,
  onClickDelete: () => void,
  onClickToEditTodo: () => void,
}

export const EditIconsArea: VFC<Props> = (props) => {
  const { finish, onChangeIsFinished, onClickDelete, onClickToEditTodo } = props

  return(
    <>
      <DefaultArea
        finish={finish}
        onChangeIsFinished={onChangeIsFinished}
        onClickDelete={onClickDelete}
        onClickSomeEvent={onClickToEditTodo}
      >
        <EditIcon 
          sx={{
            color: '#587FBA',
          }}
        />
        EDIT            
      </DefaultArea>
    </>
  )
}

export default EditIconsArea;