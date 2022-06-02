import { useSelector } from "react-redux"

import { RootState } from "../reducks/store/store";

export const useSkeletonState: () => boolean = () => {
  const skeletonState = useSelector((state: RootState) => state.skeleton.status);
  return skeletonState
}

export default useSkeletonState;