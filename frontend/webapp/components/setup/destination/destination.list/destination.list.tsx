import React from "react";
import { KeyvalText } from "@/design.system";
import { DestinationCard } from "../destination.card/destination.card";
import {
  DestinationListWrapper,
  DestinationTypeTitleWrapper,
} from "./destination.list.styled";
import { capitalizeFirstLetter } from "@/utils/functions";

export function DestinationList({
  data: { items, name },
  onItemClick,
  sectionData,
}: any) {
  function renderList() {
    return items?.map((item: any, index: number) => (
      <DestinationCard
        key={index}
        item={item}
        onClick={() => onItemClick(item)}
        focus={sectionData?.type === item?.type}
      />
    ));
  }

  return items?.length ? (
    <>
      <DestinationTypeTitleWrapper>
        <KeyvalText>{`${items?.length} ${capitalizeFirstLetter(
          name
        )}`}</KeyvalText>
      </DestinationTypeTitleWrapper>
      <DestinationListWrapper>{renderList()}</DestinationListWrapper>
    </>
  ) : null;
}
