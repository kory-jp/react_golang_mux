import { ReactNode, VFC } from "react";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props
  return(
    <>
      {children}
    </>
  )
}