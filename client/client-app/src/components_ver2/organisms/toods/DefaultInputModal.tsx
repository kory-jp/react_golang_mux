import { Modal } from "@mui/material";
import { Box } from "@mui/system";
import React, { ChangeEvent, FC } from "react";
import CloseIcon from '@mui/icons-material/Close';
import TagSelection from "../../molecules/tag/TagSelection";
import AddImageArea from "../../atoms/images/AddImageArea";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import PrimaryTextField from "../../atoms/inputs/PrimaryTextField";
import { Tags } from "../../../reducks/tags/types";
import { Todo } from "../../../reducks/todos/types";

type Props = {
  open: boolean,
  onClose: () => void,
  title: string,
  content: string,
  tags: Tags,
  importance: number | undefined,
  urgency: number | undefined,
  preview: string,
  options: Tags,
  importanceOptions: Tags,
  urgencyOptions: Tags,
  onChangeInputTitle: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onChangeInputContent: (event: React.ChangeEvent<HTMLInputElement>) => void,
  onChangeSelectTags: (event: React.SetStateAction<Tags>) => void,
  onChangeImportance:  (event: Todo) => void,
  onChangeUrgency:  (event: Todo) => void,
  onClickInputImage: (event: ChangeEvent<HTMLInputElement>) => void,
  onClickCancelImage: () => void,
  onClickNewTodo: () => void,
}

export const DefaultInputModal: FC<Props>= (props) => {
  const {
    open,
    onClose,
    title,
    content,
    tags,
    importance,
    urgency,
    preview,
    options,
    importanceOptions,
    urgencyOptions,
    onChangeInputTitle,
    onChangeInputContent,
    onChangeSelectTags,
    onChangeImportance,
    onChangeUrgency,
    onClickInputImage,
    onClickCancelImage,
    onClickNewTodo
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
                  sx={{
                    color: '#FFF',
                  }}
                />
              </Box>
            </Box>
            <Box
              className='modal__heading'
            >
              <Box
                component='h2'
                sx={{
                  color: '#FFF'
                }}
              >
                新規タスク追加
              </Box>
            </Box>
            <Box
              className='input__title'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <PrimaryTextField
                label="タイトル (必須)"
                type="text"
                value={title}
                mulitline={false}
                onChange={onChangeInputTitle}
              />
            </Box>
            <Box
              className='input__content'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <PrimaryTextField 
                label='補足'
                type='text'
                value={content}
                mulitline={true}
                rows={8}
                onChange={onChangeInputContent}
              />
            </Box>
            <Box
              className='select__tag'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <TagSelection 
                placeholder="タグを選択してください"
                options={options}
                isMulti={true}
                onChange={onChangeSelectTags}
                values={tags}
              />
            </Box>
            <Box
              className='select__importance'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <TagSelection 
                placeholder="重要度を選択してください(必須)"
                options={importanceOptions}
                isMulti={false}
                onChange={onChangeImportance}
                value={importance === undefined ? undefined : importance === 1 ? importanceOptions[0] : importanceOptions[1]}
              />
            </Box>
            <Box
              className='select__urgency'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <TagSelection 
                placeholder="緊急度を選択してください(必須)"
                options={urgencyOptions}
                isMulti={false}
                onChange={onChangeUrgency}
                value={urgency === undefined ? undefined : urgency === 1 ? urgencyOptions[0] : urgencyOptions[1]}
              />
            </Box>
            <Box
              className='input__image'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <AddImageArea 
                preview={preview}
                onClickInputImage={onClickInputImage}
                onClickCancelImage={onClickCancelImage}
              />
            </Box>
            <Box
              className='submit__button'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <PrimaryButton
                disabled={title === '' || importance === 0 || urgency === 0}
                onClick={onClickNewTodo}
              >
                Todo追加
              </PrimaryButton>
            </Box>
          </Box>
        </Box>
      </Modal>
    </>
  )
}

export default DefaultInputModal;