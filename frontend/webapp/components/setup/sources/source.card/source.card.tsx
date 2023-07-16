import {
  KeyvalCard,
  KeyvalRadioButton,
  KeyvalTag,
  KeyvalText,
} from "@/design.system";
import React from "react";
import {
  ApplicationNameWrapper,
  RadioButtonWrapper,
  SourceCardWrapper,
} from "./source.card.styled";
import Logo from "assets/logos/code-sandbox-logo.svg";
import { SETUP } from "@/utils/constants";

const KIND_COLORS = {
  deployment: "#203548",
  DaemonSet: "#033869",
};

const TEXT_STYLE = {
  textOverflow: "ellipsis",
  whiteSpace: "nowrap",
  overflow: "hidden",
};

export function SourceCard({ item, onClick, focus }: any) {
  return (
    <KeyvalCard focus={focus}>
      <RadioButtonWrapper>
        <KeyvalRadioButton onChange={onClick} value={focus} />
      </RadioButtonWrapper>
      <SourceCardWrapper onClick={onClick}>
        <Logo />
        <ApplicationNameWrapper>
          <KeyvalText size={20} weight={700} style={TEXT_STYLE}>
            {item.name}
          </KeyvalText>
        </ApplicationNameWrapper>
        <KeyvalTag title={item.kind} color={KIND_COLORS[item.kind]} />
        <KeyvalText size={14} weight={400}>
          {`${item?.instances} ${SETUP.RUNNING_INSTANCES}`}
        </KeyvalText>
      </SourceCardWrapper>
    </KeyvalCard>
  );
}
