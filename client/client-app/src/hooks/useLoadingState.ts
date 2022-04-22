import { useSelector } from "react-redux"

import { RootState } from "../reducks/store/store";

export const useLoadingState: () => boolean = () => {
  const loadingState = useSelector((state: RootState) => state.loading.status);
  return loadingState
}

export default useLoadingState;