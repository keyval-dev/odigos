import React from "react";
import { DropDown } from "@keyval-dev/design-system";

interface DropDownItem {
  id: number;
  label: string;
}
interface KeyvalDropDownProps {
  data: DropDownItem[];
  onChange: (item: DropDownItem) => void;
  width?: number;
  value?: DropDownItem | null;
}

export function KeyvalDropDown(props: KeyvalDropDownProps) {
  return <DropDown {...props} />;
}
