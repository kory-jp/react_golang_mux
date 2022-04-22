import { ReactNode, VFC } from "react";

import Header from "../organisms/layout/Header";
import { SearchSection } from "../organisms/layout/SearchSection";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props
  return(
    <>
      <Header />
      <SearchSection />
      {children}
    </>
  )
}

export default DefaultTemplate;