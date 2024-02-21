export interface ManagedSource {
  kind: string;
  name: string;
  namespace: string;
  reported_name?: string;
  languages: [
    {
      container_name: string;
      language: string;
    }
  ];
}

export interface Namespace {
  name: string;
  selected: boolean;
  totalApps: number;
}

export interface SourceConfig {
  app_instrumentation_labeled: boolean;
  instances: number;
  instrumentation_effective: boolean;
  kind: string;
  name: string;
  ns_instrumentation_labeled: boolean;
  selected?: boolean;
}

export type NamespaceConfiguration = {
  selected_all: boolean;
  future_selected: boolean;
  objects: SourceConfig[];
};

// Type for the overall structure which has "default" as a key
export type SelectedSourcesConfiguration = {
  [key: string]: NamespaceConfiguration;
};

export interface SelectedSources {
  [key: string]: {
    objects: {
      name: string;
      selected: boolean;
      kind: string;
      app_instrumentation_labeled: boolean | null;
      ns_instrumentation_labeled: boolean | null;
      instrumentation_effective: boolean | null;
      instances: number;
    };
    selected_all: boolean;
    future_selected: boolean;
  };
}
