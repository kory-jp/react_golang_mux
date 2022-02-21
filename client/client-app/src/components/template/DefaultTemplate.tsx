import { ReactNode, VFC } from "react";

import Header from "../organisms/layout/Header";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props
  return(
    <>
      <Header />
      {children}
    </>
  )
}

export default DefaultTemplate;