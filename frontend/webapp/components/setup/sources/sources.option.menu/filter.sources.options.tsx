import React from 'react';
import { DropdownWrapper } from './sources.option.menu.styled';
import { KeyvalDropDown, KeyvalSearchInput, KeyvalText } from '@/design.system';
import { SETUP } from '@/utils/constants';

const DEFAULT_DROPDOWN_VALUE = { id: 0, label: 'default' };

export function FilterSourcesOptions({
  setCurrentItem,
  data,
  searchFilter,
  setSearchFilter,
}: any) {
  function handleDropDownChange(item: any) {
    setCurrentItem({ id: item?.id, name: item.label });
  }

  return (
    <>
      <KeyvalSearchInput
        value={searchFilter}
        onChange={(e) => setSearchFilter(e.target.value)}
      />
      <DropdownWrapper>
        <KeyvalText size={14}>{SETUP.MENU.NAMESPACES}</KeyvalText>
        <KeyvalDropDown
          value={DEFAULT_DROPDOWN_VALUE}
          data={data}
          onChange={handleDropDownChange}
        />
      </DropdownWrapper>
    </>
  );
}
