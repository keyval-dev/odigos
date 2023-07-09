import React, { useState } from "react";
import Open from "@/assets/icons/expand-arrow.svg";
import {
  DropdownHeader,
  DropdownWrapper,
  DropdownBody,
  DropdownItem,
  DropdownListWrapper,
} from "./drop.down.styled";
import { KeyvalText } from "../text/text";
import { KeyvalSearchInput } from "../search.input/search.input";

interface DropDownItem {
  id: number;
  label: string;
}
interface KeyvalDropDownProps {
  data: DropDownItem[];
  onChange: (item: DropDownItem) => void;
}

const SELECTED_ITEM = "Select item";
const CONTAINER_STYLE = {
  width: "90%",
  border: "none",
  background: "transparent",
};
const SEARCH_INPUT_STYLE = { background: "transparent" };

export function KeyvalDropDown({ data, onChange }: KeyvalDropDownProps) {
  const [isOpen, setOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState<any>(data[0] || null);
  const [isHover, setHover] = useState(false);
  const [searchFilter, setSearchFilter] = useState("");

  const toggleDropdown = () => setOpen(!isOpen);

  const handleItemClick = (item: DropDownItem) => {
    onChange(item);
    setSelectedItem(item);
    setOpen(false);
  };

  function getDropdownList() {
    return searchFilter
      ? data?.filter((item: any) =>
          item?.label.toLowerCase().includes(searchFilter.toLowerCase())
        )
      : data;
  }

  return (
    <div style={{ height: 37 }}>
      <DropdownWrapper
        hover={isHover}
        onMouseEnter={() => setHover(true)}
        onMouseLeave={() => setHover(false)}
      >
        <DropdownHeader onClick={toggleDropdown}>
          {selectedItem ? selectedItem.label : SELECTED_ITEM}
          <Open className={`dropdown-arrow ${isOpen && "open"}`} />
        </DropdownHeader>
      </DropdownWrapper>
      {isOpen && (
        <DropdownBody>
          <KeyvalSearchInput
            value={searchFilter}
            onChange={(e) => setSearchFilter(e.target.value)}
            placeholder="Search"
            containerStyle={CONTAINER_STYLE}
            inputStyle={SEARCH_INPUT_STYLE}
            showClear={false}
          />
          <DropdownListWrapper>
            {getDropdownList().map((item) => (
              <DropdownItem
                key={item.id}
                onClick={(e: any) => handleItemClick(item)}
              >
                <KeyvalText>{item.label}</KeyvalText>
              </DropdownItem>
            ))}
          </DropdownListWrapper>
        </DropdownBody>
      )}
    </div>
  );
}
