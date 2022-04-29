import { ReactNode, VFC } from "react";
import Footer from "../../organisms/layout/Footer";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props
  return(
    <>
      {children}
      <Footer />
    </>
  )
}