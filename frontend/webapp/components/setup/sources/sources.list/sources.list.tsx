import React from "react";
import {
  SourcesListContainer,
  SourcesListWrapper,
  SourcesTitleWrapper,
} from "./sources.list.styled";
import { SourceCard } from "../source.card/source.card";
import { KeyvalLink, KeyvalText, KeyvalTooltip } from "@/design.system";
import { SETUP } from "@/utils/constants";

export function SourcesList({
  data,
  onItemClick,
  selectedData,
  onClearClick,
}: any) {
  function isFocus(currentCard: any) {
    const currentItem = selectedData?.objects?.filter(
      (item) => item.name === currentCard.name
    );
    return currentItem?.[0]?.selected || false;
  }

  function renderList() {
    return data?.map((item: any, index: number) => (
      <SourceCard
        key={index}
        item={item}
        onClick={() => onItemClick({ item, index })}
        focus={isFocus(item)}
      />
    ));
  }

  return (
    <SourcesListContainer>
      <SourcesTitleWrapper>
        <KeyvalText>{`${data.length} ${SETUP.APPLICATIONS}`}</KeyvalText>
        <KeyvalLink onClick={onClearClick} value={SETUP.CLEAR_SELECTION} />
      </SourcesTitleWrapper>
      <SourcesListWrapper>{renderList()}</SourcesListWrapper>
    </SourcesListContainer>
  );
}
