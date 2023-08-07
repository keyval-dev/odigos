import React from "react";
import {
  ManagedListWrapper,
  EmptyListWrapper,
  ManagedContainer,
} from "./sources.manage.styled";
import Empty from "@/assets/images/empty-list.svg";
import SourceManagedCard from "./sources.manage.card";
import { ManagedSource } from "@/types/sources";
import { KeyvalText } from "@/design.system";
import { OVERVIEW, ROUTES } from "@/utils/constants";
import { useRouter } from "next/navigation";

interface SourcesManagedListProps {
  data: ManagedSource[];
}

export function SourcesManagedList({ data = [] }: SourcesManagedListProps) {
  const router = useRouter();
  function renderSources() {
    return data.map((source: ManagedSource) => (
      <SourceManagedCard
        key={source?.name}
        item={source}
        onClick={() =>
          router.push(
            `${ROUTES.MANAGE_SOURCE}?name=${source?.name}&kind=${source?.kind}&namespace=${source?.namespace}`
          )
        }
      />
    ));
  }

  return data.length === 0 ? (
    <EmptyListWrapper>
      <Empty />
    </EmptyListWrapper>
  ) : (
    <ManagedContainer>
      <KeyvalText>{`${data.length} ${OVERVIEW.MENU.SOURCES}`}</KeyvalText>
      <br />
      <ManagedListWrapper>{renderSources()}</ManagedListWrapper>
    </ManagedContainer>
  );
}
