import { Button, Input, InputLabel } from "@mui/material";
import CameraAltIcon from '@mui/icons-material/CameraAlt';
import ClearIcon from '@mui/icons-material/Clear';
import { ChangeEvent, FC } from "react";
import { Box } from "@mui/system";

type Props ={
  preview: string,
  onClickInputImage: (event: ChangeEvent<HTMLInputElement>) => void,
  onClickCancelImage: () => void,
}

export const AddImageArea: FC<Props> = (props) => {
  const {preview, onClickInputImage, onClickCancelImage} = props
  return(
    <>
      <InputLabel 
        htmlFor="upImage"
        sx={{
          width: {
            xs: '70px',
          }
        }}
      >
        <Input
          id='upImage'
          name='image'
          type='file'
          inputProps={{accept: "image/*, .jpg, .jpeg, .png"}}
          sx={{
            display: 'none'
          }}
          onChange={onClickInputImage}
        />
        <Button variant="contained" component='span'>
          <CameraAltIcon />
        </Button>
      </InputLabel>
      {
        preview ?
        <Box>
          <Box
            textAlign='end'
          >
            <Button>
              <ClearIcon onClick={onClickCancelImage}/>
            </Button>
          </Box>
          <img 
            src={preview}
            alt='preview_img'
            width='100%'
            height='auto'
          />
        </Box> : null
      }    
    </>
  )
}

export default AddImageArea;