import { VFC } from 'react';
import Select from 'react-select'

interface Tag {
  id: number;
  value: string;
  label: string;
}

type Props = {
  placeholder: string
  isMulti: boolean
  options: Tag[]
  onChange: (event:any) => void
  value?: Tag[]
}


export const TagSelection: VFC<Props> = (props) => {
  const {placeholder, isMulti, onChange, options, value} = props

  if (isMulti) {
    return(
      <Select 
        placeholder={placeholder}
        isMulti
        options={options}
        onChange={onChange}
        value={value}
      />
    )
  } else {
    return(
      <Select 
      placeholder={placeholder}
      options={options}
      onChange={onChange}
      value={value}
    />
    )
  }
}

export default TagSelection; 