export type TMap = {
  [key: string]: any;
};

export type TGqlInput<T = null> = {
  input: T;
};

export type DropdownItem<T = string> = {
  value: T
  label: string
}

export type TSubscription = { unsubscribe?: () => void }

export type TableCellProps = {
  class: string
  label: string
}

export type FaqItem = {
  id: string
  question: string
  answer: string
}