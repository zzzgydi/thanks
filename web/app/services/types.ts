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
