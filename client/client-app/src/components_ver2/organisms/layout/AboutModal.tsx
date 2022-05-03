import { CardMedia, Divider, Grid, Link, Modal } from "@mui/material";
import { Box, maxWidth } from "@mui/system";
import { FC } from "react";
import CloseIcon from '@mui/icons-material/Close';
import Dialog from "../../../assets/images/Dialog.svg"
import taskCard from "../../../assets/images/taskCard.svg"

type Props = {
  open: boolean,
  onClose: () => void,
}

export const AboutModal: FC<Props> = (props) => {
  const {open, onClose} = props
  return (
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
                sx={{
                  marginBottom: {
                    xs: '16px'
                  }
                }}
              >
                <CloseIcon
                  fontSize="large"
                  onClick={onClose}
                  sx={{
                    color: '#FFF',
                    cursor: 'pointer',
                  }}
                />
              </Box>
            </Box>
            <Box
              className='about__heading'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <Box
                className='title'
                sx={{
                  marginBottom: {
                    xs: '24px',
                  }
                }}
              >
                <Box
                  sx={{
                    marginBottom: {
                      xs: '8px',
                    }
                  }}
                >
                  このサイトについて
                </Box>
                <Divider 
                  sx={{
                    backgroundColor: '#FFF',
                  }}
                />
              </Box>
              <Box
                className='about__introduction'
              >
                <Box
                  component='p'
                  sx={{
                    letterSpacing: '2px',
                    marginBottom: {
                      xs: '24px',
                    }
                  }}                  
                >
                  こちらのアプリはWebサイトPresident Academym様のタスク管理術を参考に作成したタスク管理効率化ツールとなります。                  
                </Box>
              </Box>
              <Box
                className='about__link'
              >
                <Link
                  underline="none"
                  href="https://president-ac.jp/blog/taskmanagement/#i"
                  sx={{
                    fontSize: {
                      xs: '16px',
                    }
                  }}
                >
                  タスク管理の上手い人がやっていることをご紹介！タスク管理に役立つツールも無料公開中
                </Link>                
              </Box>               
            </Box>
            <Box
              className='about__explanation'
            >
              <Box
                className='title'
                sx={{
                  marginBottom: {
                    xs: '16px',
                  }
                }}
              >
                <Box
                  sx={{
                    marginBottom: {
                      xs: '8px',
                    }
                  }}                
                >
                  アプリ説明
                </Box>
                <Divider 
                  sx={{
                    backgroundColor: '#FFF',
                  }}
                />                
              </Box>
              <Grid
                className="explanation__todo"
                container
                spacing={{xs: 2, md: 1}}
                sx={{
                  justifyContent: {
                    xs: 'center',
                    md: 'space-between',
                  },
                  marginBottom: {
                    xs: '40px',
                  }
                }}
              >
                <Grid
                  className='text'
                  item
                  md={6}
                >
                  <Box
                    component='p'
                    sx={{
                      letterSpacing: '2px',
                    }}
                  >
                    管理したいタスクをアプリ上部の「タスク追加」から新規登録します<br/>
                    <br/>
                    新規登録時には、「重要度」、「緊急度」を選択してタスクの優先順位を見える化することができます。
                  </Box>
                </Grid>
                <Grid
                  className='dialog'
                  item
                  md={5}
                >
                  <CardMedia
                    component="img"
                    image={Dialog}
                    sx={{
                      height : {
                        xs: 'auto',
                      }
                    }}
                  />                  
                </Grid>
              </Grid>
              <Grid
                className="explanation__tc"
                container
                spacing={{xs:2, md:1}}
                sx={{
                  justifyContent: {
                    xs: 'center',
                    md: 'space-between',
                  }
                }}
              >
                <Grid
                  className="text"
                  item
                  md={6}
                >
                  <Box>
                    タスクを追加後は、タスクカード機能を用いてタスクの細分化、目的と作業内容を記録することでタスクの明確化を図りましょう。
                  </Box>
                </Grid>
                <Grid
                  className="img"
                  item
                  md={5}
                >
                  <CardMedia
                    component="img"
                    image={taskCard}
                    sx={{
                      height : {
                        xs: 'auto',
                      }
                    }}
                  />                     
                </Grid>
              </Grid>
            </Box>
          </Box>
        </Box>
      </Modal>       
    </>
  )
}

export default AboutModal;