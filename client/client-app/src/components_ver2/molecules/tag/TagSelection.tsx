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
  values?: Tag[]
  value?: Tag
}

const customStyles = {
  control: (base: any) => ({
    ...base,
    background: "#605D5D",
    color: '#FFF',
    fontSize: '14px',
    minWidth: "227px",
  }),
  option: (base: any) => ({
    color: '#000',
    cursor: 'pointer',
  }),
  placeholder: (defaultStyles: any) => {
    return {
      ...defaultStyles,
      color: "#FFF"
    }
  },
  singleValue: (base: any) => ({
    ...base,
    color: "#FFF",
  })
};

export const TagSelection: VFC<Props> = (props) => {
  const {placeholder, isMulti, onChange, options, values, value} = props

  if (isMulti) {
    return(
      <Select 
        placeholder={placeholder}
        isMulti
        options={options}
        onChange={onChange}
        value={values}
        styles={customStyles}
      />
    )
  } else {
    return(
      <Select 
        placeholder={placeholder}
        options={options}
        onChange={onChange}
        value={value}
        styles={customStyles}
      />
    )
  }
}

export default TagSelection; 