import React, { useState } from 'react';
import { ActionState, ActionsType } from '@/types';
import DeleteAttributeYaml from './delete.attribute.yaml';
import { KeyvalText } from '@/design.system';
import styled from 'styled-components';
import AddClusterInfoYaml from './add.cluster.info.yaml';

import theme from '@/styles/palette';
import { Check, YamlIcon } from '@/assets/icons/app';

const CodeBlockWrapper = styled.p`
  display: flex;
  align-items: center;
  font-family: Inter;
  color: ${theme.text.light_grey};
  a {
    color: ${theme.text.secondary};
    text-decoration: none;
    cursor: pointer;
  }
`;
const TitleWrapper = styled.div`
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 8px;
`;

interface ActionsYamlProps {
  data: ActionState;
  onChange: (key: string, value: any) => void;
}

export function ActionsYaml({ data, onChange }: ActionsYamlProps) {
  const [copied, setCopied] = useState(false);
  const [echoCommand, setEchoCommand] = useState('');

  function renderYamlEditor() {
    switch (data.type) {
      case ActionsType.DELETE_ATTRIBUTES:
        return (
          <DeleteAttributeYaml
            data={data}
            onChange={onChange}
            setEchoCommand={setEchoCommand}
          />
        );
      case ActionsType.ADD_CLUSTER_INFO:
        return (
          <AddClusterInfoYaml
            data={data}
            onChange={onChange}
            setEchoCommand={setEchoCommand}
          />
        );
      default:
        return <></>;
    }
  }

  function handleCopy() {
    navigator.clipboard.writeText(echoCommand);
    setCopied(true);
    setTimeout(() => {
      setCopied(false);
    }, 3000);
  }

  return (
    <div>
      <TitleWrapper>
        <YamlIcon style={{ width: 20, height: 20 }} />
        <KeyvalText
          weight={600}
        >{`YAML Preview - ${data.type.toLowerCase()}.actions.odigos.io`}</KeyvalText>
      </TitleWrapper>

      <TitleWrapper>
        <KeyvalText size={14}>
          This is the YAML representation of the action you are creating.
        </KeyvalText>
      </TitleWrapper>
      {renderYamlEditor()}
      <div style={{ width: 600, overflowX: 'hidden' }}>
        <CodeBlockWrapper>
          {copied ? (
            <Check style={{ width: 18, height: 12 }} />
          ) : (
            <YamlIcon style={{ width: 18, height: 18 }} />
          )}
          <a style={{ margin: '0 4px' }} onClick={handleCopy}>
            Click here
          </a>
          to copy as kubectl command.
        </CodeBlockWrapper>
      </div>
    </div>
  );
}
