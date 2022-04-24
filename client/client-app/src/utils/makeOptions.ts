export type Option = {
  id: number;
  value: string;
  label: string;
};

export type Options = Option[];

export const makeOptions: () => {
  importanceOptions: Options;
  urgencyOptions: Options;
} = () => {
  const importanceOptions: Options = [
    { id: 1, value: "Importance", label: "重要" },
    { id: 2, value: "NotImportance", label: "重要ではない" }
  ];

  const urgencyOptions: Options = [
    { id: 1, value: "Urgency", label: "緊急" },
    { id: 2, value: "NotUrgency", label: "緊急ではない" }
  ];

  return { importanceOptions, urgencyOptions };
};
