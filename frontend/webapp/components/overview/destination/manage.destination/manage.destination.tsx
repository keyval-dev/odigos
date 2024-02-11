import React, { useMemo } from 'react';
import { styled } from 'styled-components';
import { Back } from '@/assets/icons/overview';
import { CreateConnectionForm, QuickHelp } from '@/components/setup';
import { KeyvalText } from '@/design.system';
import { SETUP } from '@/utils/constants';
import { ManageDestinationHeader } from '../manage.destination.header/manage.destination.header';
import { DestinationType } from '@/types/destinations';
import FormDangerZone from './form.danger.zone';

interface ManageDestinationProps {
  destinationType: DestinationType;
  selectedDestination: any;
  onBackClick?: () => void;
  onSubmit: (data: any) => void;
  onDelete?: () => void;
}

const BackButtonWrapper = styled.div`
  display: flex;
  align-items: center;
  cursor: pointer;
  p {
    cursor: pointer !important;
  }
`;

const CreateConnectionWrapper = styled.div`
  display: flex;
  gap: 10vw;
`;

export function ManageDestination({
  destinationType,
  selectedDestination,
  onBackClick,
  onSubmit,
  onDelete,
}: ManageDestinationProps) {
  const videoList = useMemo(
    () =>
      destinationType?.fields
        ?.filter((field) => field?.video_url)
        ?.map((field) => ({
          name: field.display_name,
          src: field.video_url,
          thumbnail_url: field.thumbnail_url,
        })),
    [destinationType]
  );

  console.log({ field: destinationType?.fields });
  console.log({ dynamicFieldsValues: selectedDestination?.fields });

  return (
    <>
      {onBackClick && (
        <BackButtonWrapper onClick={onBackClick}>
          <Back width={14} />
          <KeyvalText size={14}>{SETUP.BACK}</KeyvalText>
        </BackButtonWrapper>
      )}
      <ManageDestinationHeader data={selectedDestination} />
      <CreateConnectionWrapper>
        <div>
          <CreateConnectionForm
            fields={destinationType?.fields}
            destinationNameValue={selectedDestination?.name}
            dynamicFieldsValues={selectedDestination?.fields}
            checkboxValues={selectedDestination?.signals}
            supportedSignals={selectedDestination?.supported_signals}
            onSubmit={(data) => onSubmit(data)}
          />
          {onDelete && (
            <FormDangerZone onDelete={onDelete} data={selectedDestination} />
          )}
        </div>
        {videoList?.length > 0 && <QuickHelp data={videoList} />}
      </CreateConnectionWrapper>
    </>
  );
}
