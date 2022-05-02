import { Modal } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";
import CloseIcon from '@mui/icons-material/Close';
import PrimaryTextField from "../../atoms/inputs/PrimaryTextField";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";

type Props = {
  open: boolean,
  onClose: () => void,
  title: string,
  purpose: string,
  content: string,
  memo: string,
  onChangeInputTitle: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onChangeInputPurpose: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onChangeInputContent: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onChangeInputMemo: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onClickSubmitTC: () => void
}

export const DefaultInputTCModal: FC<Props> = (props) => {
  const {
    open, 
    onClose, 
    title, 
    purpose, 
    content, 
    memo,
    onChangeInputTitle,
    onChangeInputPurpose,
    onChangeInputContent,
    onChangeInputMemo,
    onClickSubmitTC
  } = props

  return(
    <>
      <Modal
        open={open}
        onClose={onClose}
        sx={{
          overflow: 'scroll',
        }}
      >
        <Box
          className='createTodoModal'
          bgcolor="#2D2A2A"
          sx={{
            marginX: 'auto',
            marginTop: '5%',
            width: {
              xs: '90%',
              sm: '70%',
              md: '60%',
              lg: '50%',
            },
            borderRadius: '10px',
          }}
        >
          <Box
            className='modal__inner'
            sx={{
              padding: {
                xs: '16px',
              }
            }}
          >
            <Box
              className='close'
              textAlign='end'
            >
              <Box
                className='close__button'
                onClick={onClose}
              >
                <CloseIcon
                  fontSize="large"
                />
              </Box>
            </Box>
            <Box
              className='create_tc_heading'
            >
              <Box
                component='h2'
                sx={{
                  marginBottom: {
                    xs: '40px',
                  }
                }}
              >
                新規タスクカード追加
              </Box>
            </Box>
            <Box
              className='tc_title'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <Box>
                <PrimaryTextField 
                  label='タイトル'
                  value={title}
                  type='text'
                  mulitline={false}
                  onChange={onChangeInputTitle}
                />
              </Box>
            </Box>
            <Box
              className='tc__purpose'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}              
            >
              <Box >
                目的
              </Box>
              <Box>
                <PrimaryTextField 
                  label='なぜこのタスクをする必要があるのか'
                  value={purpose}
                  type='text'
                  mulitline={true}
                  rows={6}
                  onChange={onChangeInputPurpose}
                />
              </Box>
            </Box>
            <Box
              className='tc__content'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}   
            >
              <Box>
                作業内容
              </Box>
              <Box>
                <PrimaryTextField 
                  label="具体的にどのような作業をするのか"
                  value={content}
                  type='text'
                  mulitline={true}
                  rows={6}
                  onChange={onChangeInputContent}
                />
              </Box>
            </Box>
            <Box
              className='tc__memo'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }} 
            >
              <Box>
                メモ
              </Box>
              <Box>
                <PrimaryTextField 
                  label=''
                  value={memo}
                  type='text'
                  mulitline={true}
                  rows={6}
                  onChange={onChangeInputMemo}
                />
              </Box>
            </Box>
            <Box
              className='tc__submit'
            >
              <Box>
                <PrimaryButton
                  onClick={onClickSubmitTC}
                >
                  タスクカード追加
                </PrimaryButton>
              </Box>
            </Box>
          </Box>
        </Box>
      </Modal>          
    </>
  )
}

export default DefaultInputTCModal;