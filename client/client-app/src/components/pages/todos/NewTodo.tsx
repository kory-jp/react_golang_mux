import CameraAltIcon from '@mui/icons-material/CameraAlt';
import ClearIcon from '@mui/icons-material/Clear';
import { Button, Divider, FormControl, Input, InputLabel, Paper, Stack, TextField} from "@mui/material";
import { Box } from "@mui/system";
import React, { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { createTodo } from "../../../reducks/todos/operations";
import TagSelection from '../../molecules/tag/TagSelction';
import { indexTags } from '../../../reducks/tags/operations';
import { RootState } from '../../../reducks/store/store';
import { Tags } from '../../../reducks/tags/types';

export const NewTodo: FC = () => {
  const dispatch = useDispatch()
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [image, setImage] = useState<File>()
  const [preview, setPreview] =useState('')
  const [tags, setTags] = useState<Tags>([])

  useEffect(() => {
    dispatch(indexTags())
  }, [dispatch])

  const options = useSelector((state: RootState) => state.tags)
  
  const inputTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])
  
  const inputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])

  const onChangeSelectTags = useCallback((event: React.SetStateAction<Tags>) => {
    setTags(event)
  }, [setTags])
  
  const previewImage = useCallback((event) => {
    const imageFile = event.target.files[0];
    setPreview(window.URL.createObjectURL(imageFile))
  },[])
  
  const inputImage = useCallback((event)=> {
    const file = event.target.files[0]
    setImage(file)
    previewImage(event)
  }, [setImage, previewImage])

  const onClickCancelImage = useCallback(() => {
    setImage(undefined)
    setPreview('')
  }, [])


  const createFormData = useCallback(() => {
    const formData = new FormData()
    formData.append('title', title)
    formData.append('content', content)
    if (image) formData.append('image', image)
    for(let i in tags) {
      let tagId = String(tags[i].id)
      formData.append('tagIds', tagId)
    }
    return formData
  }, [title, content, image, tags])


  const formData = createFormData()
  const onClickNewTodo = useCallback(() => {
    dispatch(createTodo(formData))
  }, [formData])

  return(
    <>
      <Paper
        sx={{
          padding: {
            xs: '10px',
            md: '30px'
          },
          marginX: {
            xs: '0px',
            sm: '80px',
            md: '150px'
          }
        }}
      >
        <Box
          fontWeight='bold'
          marginBottom='10px'
          sx={{
            fontSize: {
              sx: '20px',
              md: '35px'
            }
          }}
        >
          新規TODO追加
        </Box>
        <Divider />
        <Box
          sx={{
            margin: {
              sx: '10px',
              md: '20px'
            }
          }}
        >
          <Stack spacing={3}>
            <FormControl fullWidth>
              <TextField 
                label='Ttile'
                type='text'
                value={title}
                variant="standard"
                onChange={inputTitle}
              />
            </FormControl>
            <FormControl fullWidth>
              <TextField 
                label='Memo'
                multiline
                type='text'
                value={content}
                variant="standard"
                rows={4}
                onChange={inputContent}
              />
            </FormControl>
            {/* --- タグ選択 --- */}
            <FormControl fullWidth>
              <TagSelection isMulti={true} options={options} onChange={onChangeSelectTags}/>
            </FormControl>
            {/* <FormControl> ----削除------ */}
            <InputLabel htmlFor="upImage">
              <Input 
                id='upImage'
                name='image'
                type='file'
                inputProps={{accept: "image/*, .jpg, .jpeg, .png"}}
                onChange={inputImage}
                sx={{
                  display: 'none'
                }}
              />
              <Button variant="contained" component='span'>
                <CameraAltIcon />
              </Button>
            </InputLabel>
            {/* </FormControl> */}
            {
              preview ?
              <Box>
                <Button>
                  <ClearIcon onClick={onClickCancelImage}/>
                </Button>
                <img 
                  src={preview}
                  alt='preview_img'
                  width='100%'
                  height='auto'
                />
              </Box> : null
            }
            <PrimaryButton
              disabled={title === ''}
              onClick={onClickNewTodo}
            >
              追加
            </PrimaryButton>
          </Stack>
        </Box>
      </Paper>
    </>
  )
}

export default NewTodo;