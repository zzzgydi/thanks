export interface IThkContributor {
  login: string;
  id: number;
  total: number;
  repos: IThkContributorRepo[];
}

export interface IThkContributorRepo {
  repo: string;
  score: number;
}

export interface IThkTaskResponse {
  id: string;
  lang: string;
  list: IThkContributor[];
  total: number;
  offset: number;
  created_at: string;
  updated_at: string;
}
