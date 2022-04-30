import { Button, Card, CardActions, CardContent, CardMedia, Checkbox, Chip, FormControlLabel, Grid, Typography } from "@mui/material";
import { Dispatch, FC, SetStateAction, useCallback, useEffect, useState } from "react";
import DeleteIcon from '@mui/icons-material/Delete';
import StickyNote2Icon from '@mui/icons-material/StickyNote2';
import { useDispatch } from "react-redux";

import DefaultImage from "../../../assets/images/DefaultImage.jpg"
import { Todo } from "../../../reducks/todos/types";
import { deleteTodoInIndex, updateIsFinished } from "../../../reducks/todos/operations";
import { push } from "connected-react-router";
import { Tags } from "../../../reducks/tags/types";
import { Box } from "@mui/system";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import { PrimaryChip } from "../../atoms/chip/PrimaryChip";

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

  useEffect(() => {
    setFinish(todo.isFinished)
  }, [])

  const onclickToShowTodo = useCallback(() => {
    dispatch(push(`/todo/show/${todo.id}`))
  }, [])

  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(todo.id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(todo.id, true))
    }
  }, [todo, finish])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodoInIndex(todo.id, setSumPage, queryPage))
  }, [todo])

  const onClickToSearchTagTodo = useCallback((tagId: number) => {
    dispatch(push(`/todo/tag/${tagId}`))
  },[])

  return(
    <>
      <Card
        sx={{
          transition: '0.7s',
          bgcolor: finish? 'text.disabled' : '#2D2A2A',
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
          <Grid
            container
          >
            <Grid item xs={5}>
              <FormControlLabel
                control={<Checkbox 
                            checked={finish}
                            value={finish}
                            onChange={onChangeIsFinished}
                            sx={{
                              color: '#587FBA',
                            }}
                          />} 
                sx={{
                  color: '#FFF',
                  fontSize: '8px',
                  marginBottom: '3px',
                  fontFamily: 'Merriweather, serif',
                }}
                label="finish"
              />
            </Grid>
            <Grid 
              item 
              xs={7}
              sx={{
                display: 'flex',
              }}
            >
              <Box
                marginRight='8px'
              >
                <Button
                  onClick={onClickDelete}
                  sx={{
                    color: '#FFF',
                    fontSize: '8px',
                    marginBottom: '3px',
                    fontFamily: 'Merriweather, serif',
                  }}
                >
                  <DeleteIcon 
                    sx={{
                      color: '#587FBA',
                    }}
                  />
                  Delete
                </Button>
              </Box>
              <Box>
                <Button
                  onClick={onclickToShowTodo}
                  sx={{
                    color: '#FFF',
                    fontSize: '8px',
                    fontFamily: 'Merriweather, serif',
                  }}
                >
                  <StickyNote2Icon 
                    sx={{
                      color: '#587FBA',
                    }}
                  />
                  more
                </Button>
              </Box>
            </Grid>
          </Grid>
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
          <Grid
            container
            sx={{
              minHeight: '40px',
            }}
          >
            {
              tags != null && (
                <>
                  {
                    tags.map(tag => (
                      <Grid
                        key={tag.id}
                        marginLeft='10px'
                      >
                        <PrimaryChip 
                          label={tag.label}
                          onClick={() => onClickToSearchTagTodo(tag.id)}
                        />                      
                      </Grid>
                    ))
                  }                
                </>
              )
            }
          </Grid>
      </Card>      
    </>
  )
}

export default TodoCard;