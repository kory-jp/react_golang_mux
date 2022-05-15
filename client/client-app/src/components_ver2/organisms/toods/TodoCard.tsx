import { Card, CardActions, CardContent, CardMedia, Typography } from "@mui/material";
import { Dispatch, FC, SetStateAction, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";

import DefaultImage from "../../../assets/images/DefaultImage.jpg"
import { Todo } from "../../../reducks/todos/types";
import { deleteTodoInIndex, updateIsFinished } from "../../../reducks/todos/operations";
import { push } from "connected-react-router";
import { Tags } from "../../../reducks/tags/types";
import { Box } from "@mui/system";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import MoreIconsArea from "../../molecules/iconsArea/MoreIconsArea";
import TagSection from "../../molecules/tag/TagSection";
import useReturnTop from "../../../hooks/useReturnTop";

type Props = {
  todo: Todo,
  setSumPage :Dispatch<SetStateAction<number>>,
  queryPage: number
}

export const TodoCard: FC<Props> = (props) => {
  const {todo, setSumPage, queryPage} = props
  const tags: Tags | null = todo.tags ? todo.tags : null
  const dispatch = useDispatch()
  const [finish, setFinish] = useState(false)
  const imagePath = process.env.REACT_APP_API_URL + `img/${todo.imagePath}`
  const returnTop = useReturnTop()

  useEffect(() => {
    setFinish(todo.isFinished)
  }, [todo.isFinished])

  const onclickToShowTodo = useCallback(() => {
    dispatch(push(`/todo/show/${todo.id}`))
    returnTop()
  }, [dispatch, returnTop, todo.id])

  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(todo.id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(todo.id, true))
    }
  }, [dispatch, todo, finish])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodoInIndex(todo.id, setSumPage, queryPage))
  }, [dispatch, setSumPage, queryPage, todo])

  const onClickToSearchTagTodo = useCallback((tagId: number) => {
    dispatch(push(`/todo/search?tagId=${tagId}&importance=0&urgency=0&page=1`))
    returnTop()
  },[dispatch, returnTop])

  return(
    <>
      <Card
        sx={{
          transition: '0.7s',
          bgcolor: finish? '#464141' : '#2D2A2A',
          borderRadius: '10px',
          width: {
            md: '360px',
            xl: '400px',
          }
        }}
      >
        {/* ----- 画像セクション ----- */}
        <CardMedia
          component="img"
          image={todo.imagePath? imagePath : DefaultImage}
          sx={{
            height : '200px',
            transition: '0.7s',
            filter: finish? 'grayscale(100%)' : '',
            '&:hover': {
              cursor: 'pointer'
            }
          }}
          onClick={onclickToShowTodo}
        />
        {/* ----- タイトルセクション ----- */}
        <CardContent>
          <Typography
            sx={{
              fontFamily: 'Noto Serif JP, serif',
              color: '#FFF',
              '&:hover': {
                cursor: 'pointer'
              }
            }}
            onClick={onclickToShowTodo}
          >
            {todo.title}
          </Typography>
        </CardContent>
        {/* ----- 編集セクション ----- */}
        <CardActions
          sx={{
            height: '32px',
            paddingTop: '0px',
            paddingBottom: '16px',
            paddingX: '16px',
          }}          
        >
          <MoreIconsArea 
            finish={finish}
            onChangeIsFinished={onChangeIsFinished}
            onClickDelete={onClickDelete}
            onClickMoreInfo={onclickToShowTodo}
          />
        </CardActions>
        {/* --- timezone --- */}
        <CardContent
          sx={{
            display: 'flex',
            justifyContent: 'flex-end',
            paddingTop: '0px',
            paddingBottom: '16px',
          }}
        >
          <Box
            color='#FFF'
          >
            {handleToDateFormat(todo.createdAt)}
          </Box>
        </CardContent>
        <Box
          sx={{
            paddingX: {
              xs: '16px',
            }
          }}
        >
          <TagSection 
            tags={tags}
            onClickToSearchTagTodo={onClickToSearchTagTodo}
          />
        </Box>
      </Card>      
    </>
  )
}

export default TodoCard;