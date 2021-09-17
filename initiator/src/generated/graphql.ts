export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  bigint: any;
  numeric: any;
  timestamptz: any;
};

/** expression to compare columns of type Int. All fields are combined with logical 'AND'. */
export type Int_Comparison_Exp = {
  _eq?: Maybe<Scalars['Int']>;
  _gt?: Maybe<Scalars['Int']>;
  _gte?: Maybe<Scalars['Int']>;
  _in?: Maybe<Array<Scalars['Int']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['Int']>;
  _lte?: Maybe<Scalars['Int']>;
  _neq?: Maybe<Scalars['Int']>;
  _nin?: Maybe<Array<Scalars['Int']>>;
};

/** expression to compare columns of type String. All fields are combined with logical 'AND'. */
export type String_Comparison_Exp = {
  _eq?: Maybe<Scalars['String']>;
  _gt?: Maybe<Scalars['String']>;
  _gte?: Maybe<Scalars['String']>;
  _ilike?: Maybe<Scalars['String']>;
  _in?: Maybe<Array<Scalars['String']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _like?: Maybe<Scalars['String']>;
  _lt?: Maybe<Scalars['String']>;
  _lte?: Maybe<Scalars['String']>;
  _neq?: Maybe<Scalars['String']>;
  _nilike?: Maybe<Scalars['String']>;
  _nin?: Maybe<Array<Scalars['String']>>;
  _nlike?: Maybe<Scalars['String']>;
  _nsimilar?: Maybe<Scalars['String']>;
  _similar?: Maybe<Scalars['String']>;
};

/** expression to compare columns of type bigint. All fields are combined with logical 'AND'. */
export type Bigint_Comparison_Exp = {
  _eq?: Maybe<Scalars['bigint']>;
  _gt?: Maybe<Scalars['bigint']>;
  _gte?: Maybe<Scalars['bigint']>;
  _in?: Maybe<Array<Scalars['bigint']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['bigint']>;
  _lte?: Maybe<Scalars['bigint']>;
  _neq?: Maybe<Scalars['bigint']>;
  _nin?: Maybe<Array<Scalars['bigint']>>;
};

/** columns and relationships of "chainlink_configs" */
export type Chainlink_Configs = {
  __typename?: 'chainlink_configs';
  chainId: Scalars['bigint'];
  confirmNewAssetJobId: Scalars['String'];
  cookie: Scalars['String'];
  debridgeAddr: Scalars['String'];
  eiChainlinkUrl: Scalars['String'];
  eiCiAccesskey: Scalars['String'];
  eiCiSecret: Scalars['String'];
  eiIcAccesskey: Scalars['String'];
  eiIcSecret: Scalars['String'];
  network: Scalars['String'];
  provider: Scalars['String'];
  submitJobId: Scalars['String'];
  submitManyJobId: Scalars['String'];
};

/** aggregated selection of "chainlink_configs" */
export type Chainlink_Configs_Aggregate = {
  __typename?: 'chainlink_configs_aggregate';
  aggregate?: Maybe<Chainlink_Configs_Aggregate_Fields>;
  nodes: Array<Chainlink_Configs>;
};

/** aggregate fields of "chainlink_configs" */
export type Chainlink_Configs_Aggregate_Fields = {
  __typename?: 'chainlink_configs_aggregate_fields';
  avg?: Maybe<Chainlink_Configs_Avg_Fields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<Chainlink_Configs_Max_Fields>;
  min?: Maybe<Chainlink_Configs_Min_Fields>;
  stddev?: Maybe<Chainlink_Configs_Stddev_Fields>;
  stddev_pop?: Maybe<Chainlink_Configs_Stddev_Pop_Fields>;
  stddev_samp?: Maybe<Chainlink_Configs_Stddev_Samp_Fields>;
  sum?: Maybe<Chainlink_Configs_Sum_Fields>;
  var_pop?: Maybe<Chainlink_Configs_Var_Pop_Fields>;
  var_samp?: Maybe<Chainlink_Configs_Var_Samp_Fields>;
  variance?: Maybe<Chainlink_Configs_Variance_Fields>;
};


/** aggregate fields of "chainlink_configs" */
export type Chainlink_Configs_Aggregate_FieldsCountArgs = {
  columns?: Maybe<Array<Chainlink_Configs_Select_Column>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "chainlink_configs" */
export type Chainlink_Configs_Aggregate_Order_By = {
  avg?: Maybe<Chainlink_Configs_Avg_Order_By>;
  count?: Maybe<Order_By>;
  max?: Maybe<Chainlink_Configs_Max_Order_By>;
  min?: Maybe<Chainlink_Configs_Min_Order_By>;
  stddev?: Maybe<Chainlink_Configs_Stddev_Order_By>;
  stddev_pop?: Maybe<Chainlink_Configs_Stddev_Pop_Order_By>;
  stddev_samp?: Maybe<Chainlink_Configs_Stddev_Samp_Order_By>;
  sum?: Maybe<Chainlink_Configs_Sum_Order_By>;
  var_pop?: Maybe<Chainlink_Configs_Var_Pop_Order_By>;
  var_samp?: Maybe<Chainlink_Configs_Var_Samp_Order_By>;
  variance?: Maybe<Chainlink_Configs_Variance_Order_By>;
};

/** input type for inserting array relation for remote table "chainlink_configs" */
export type Chainlink_Configs_Arr_Rel_Insert_Input = {
  data: Array<Chainlink_Configs_Insert_Input>;
  on_conflict?: Maybe<Chainlink_Configs_On_Conflict>;
};

/** aggregate avg on columns */
export type Chainlink_Configs_Avg_Fields = {
  __typename?: 'chainlink_configs_avg_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Avg_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** Boolean expression to filter rows from the table "chainlink_configs". All fields are combined with a logical 'AND'. */
export type Chainlink_Configs_Bool_Exp = {
  _and?: Maybe<Array<Maybe<Chainlink_Configs_Bool_Exp>>>;
  _not?: Maybe<Chainlink_Configs_Bool_Exp>;
  _or?: Maybe<Array<Maybe<Chainlink_Configs_Bool_Exp>>>;
  chainId?: Maybe<Bigint_Comparison_Exp>;
  confirmNewAssetJobId?: Maybe<String_Comparison_Exp>;
  cookie?: Maybe<String_Comparison_Exp>;
  debridgeAddr?: Maybe<String_Comparison_Exp>;
  eiChainlinkUrl?: Maybe<String_Comparison_Exp>;
  eiCiAccesskey?: Maybe<String_Comparison_Exp>;
  eiCiSecret?: Maybe<String_Comparison_Exp>;
  eiIcAccesskey?: Maybe<String_Comparison_Exp>;
  eiIcSecret?: Maybe<String_Comparison_Exp>;
  network?: Maybe<String_Comparison_Exp>;
  provider?: Maybe<String_Comparison_Exp>;
  submitJobId?: Maybe<String_Comparison_Exp>;
  submitManyJobId?: Maybe<String_Comparison_Exp>;
};

/** unique or primary key constraints on table "chainlink_configs" */
export enum Chainlink_Configs_Constraint {
  /** unique or primary key constraint */
  ChainlinkConfigsPkey = 'chainlink_configs_pkey'
}

/** input type for incrementing integer column in table "chainlink_configs" */
export type Chainlink_Configs_Inc_Input = {
  chainId?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "chainlink_configs" */
export type Chainlink_Configs_Insert_Input = {
  chainId?: Maybe<Scalars['bigint']>;
  confirmNewAssetJobId?: Maybe<Scalars['String']>;
  cookie?: Maybe<Scalars['String']>;
  debridgeAddr?: Maybe<Scalars['String']>;
  eiChainlinkUrl?: Maybe<Scalars['String']>;
  eiCiAccesskey?: Maybe<Scalars['String']>;
  eiCiSecret?: Maybe<Scalars['String']>;
  eiIcAccesskey?: Maybe<Scalars['String']>;
  eiIcSecret?: Maybe<Scalars['String']>;
  network?: Maybe<Scalars['String']>;
  provider?: Maybe<Scalars['String']>;
  submitJobId?: Maybe<Scalars['String']>;
  submitManyJobId?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type Chainlink_Configs_Max_Fields = {
  __typename?: 'chainlink_configs_max_fields';
  chainId?: Maybe<Scalars['bigint']>;
  confirmNewAssetJobId?: Maybe<Scalars['String']>;
  cookie?: Maybe<Scalars['String']>;
  debridgeAddr?: Maybe<Scalars['String']>;
  eiChainlinkUrl?: Maybe<Scalars['String']>;
  eiCiAccesskey?: Maybe<Scalars['String']>;
  eiCiSecret?: Maybe<Scalars['String']>;
  eiIcAccesskey?: Maybe<Scalars['String']>;
  eiIcSecret?: Maybe<Scalars['String']>;
  network?: Maybe<Scalars['String']>;
  provider?: Maybe<Scalars['String']>;
  submitJobId?: Maybe<Scalars['String']>;
  submitManyJobId?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Max_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmNewAssetJobId?: Maybe<Order_By>;
  cookie?: Maybe<Order_By>;
  debridgeAddr?: Maybe<Order_By>;
  eiChainlinkUrl?: Maybe<Order_By>;
  eiCiAccesskey?: Maybe<Order_By>;
  eiCiSecret?: Maybe<Order_By>;
  eiIcAccesskey?: Maybe<Order_By>;
  eiIcSecret?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
  provider?: Maybe<Order_By>;
  submitJobId?: Maybe<Order_By>;
  submitManyJobId?: Maybe<Order_By>;
};

/** aggregate min on columns */
export type Chainlink_Configs_Min_Fields = {
  __typename?: 'chainlink_configs_min_fields';
  chainId?: Maybe<Scalars['bigint']>;
  confirmNewAssetJobId?: Maybe<Scalars['String']>;
  cookie?: Maybe<Scalars['String']>;
  debridgeAddr?: Maybe<Scalars['String']>;
  eiChainlinkUrl?: Maybe<Scalars['String']>;
  eiCiAccesskey?: Maybe<Scalars['String']>;
  eiCiSecret?: Maybe<Scalars['String']>;
  eiIcAccesskey?: Maybe<Scalars['String']>;
  eiIcSecret?: Maybe<Scalars['String']>;
  network?: Maybe<Scalars['String']>;
  provider?: Maybe<Scalars['String']>;
  submitJobId?: Maybe<Scalars['String']>;
  submitManyJobId?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Min_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmNewAssetJobId?: Maybe<Order_By>;
  cookie?: Maybe<Order_By>;
  debridgeAddr?: Maybe<Order_By>;
  eiChainlinkUrl?: Maybe<Order_By>;
  eiCiAccesskey?: Maybe<Order_By>;
  eiCiSecret?: Maybe<Order_By>;
  eiIcAccesskey?: Maybe<Order_By>;
  eiIcSecret?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
  provider?: Maybe<Order_By>;
  submitJobId?: Maybe<Order_By>;
  submitManyJobId?: Maybe<Order_By>;
};

/** response of any mutation on the table "chainlink_configs" */
export type Chainlink_Configs_Mutation_Response = {
  __typename?: 'chainlink_configs_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Chainlink_Configs>;
};

/** input type for inserting object relation for remote table "chainlink_configs" */
export type Chainlink_Configs_Obj_Rel_Insert_Input = {
  data: Chainlink_Configs_Insert_Input;
  on_conflict?: Maybe<Chainlink_Configs_On_Conflict>;
};

/** on conflict condition type for table "chainlink_configs" */
export type Chainlink_Configs_On_Conflict = {
  constraint: Chainlink_Configs_Constraint;
  update_columns: Array<Chainlink_Configs_Update_Column>;
  where?: Maybe<Chainlink_Configs_Bool_Exp>;
};

/** ordering options when selecting data from "chainlink_configs" */
export type Chainlink_Configs_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmNewAssetJobId?: Maybe<Order_By>;
  cookie?: Maybe<Order_By>;
  debridgeAddr?: Maybe<Order_By>;
  eiChainlinkUrl?: Maybe<Order_By>;
  eiCiAccesskey?: Maybe<Order_By>;
  eiCiSecret?: Maybe<Order_By>;
  eiIcAccesskey?: Maybe<Order_By>;
  eiIcSecret?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
  provider?: Maybe<Order_By>;
  submitJobId?: Maybe<Order_By>;
  submitManyJobId?: Maybe<Order_By>;
};

/** primary key columns input for table: "chainlink_configs" */
export type Chainlink_Configs_Pk_Columns_Input = {
  chainId: Scalars['bigint'];
};

/** select columns of table "chainlink_configs" */
export enum Chainlink_Configs_Select_Column {
  /** column name */
  ChainId = 'chainId',
  /** column name */
  ConfirmNewAssetJobId = 'confirmNewAssetJobId',
  /** column name */
  Cookie = 'cookie',
  /** column name */
  DebridgeAddr = 'debridgeAddr',
  /** column name */
  EiChainlinkUrl = 'eiChainlinkUrl',
  /** column name */
  EiCiAccesskey = 'eiCiAccesskey',
  /** column name */
  EiCiSecret = 'eiCiSecret',
  /** column name */
  EiIcAccesskey = 'eiIcAccesskey',
  /** column name */
  EiIcSecret = 'eiIcSecret',
  /** column name */
  Network = 'network',
  /** column name */
  Provider = 'provider',
  /** column name */
  SubmitJobId = 'submitJobId',
  /** column name */
  SubmitManyJobId = 'submitManyJobId'
}

/** input type for updating data in table "chainlink_configs" */
export type Chainlink_Configs_Set_Input = {
  chainId?: Maybe<Scalars['bigint']>;
  confirmNewAssetJobId?: Maybe<Scalars['String']>;
  cookie?: Maybe<Scalars['String']>;
  debridgeAddr?: Maybe<Scalars['String']>;
  eiChainlinkUrl?: Maybe<Scalars['String']>;
  eiCiAccesskey?: Maybe<Scalars['String']>;
  eiCiSecret?: Maybe<Scalars['String']>;
  eiIcAccesskey?: Maybe<Scalars['String']>;
  eiIcSecret?: Maybe<Scalars['String']>;
  network?: Maybe<Scalars['String']>;
  provider?: Maybe<Scalars['String']>;
  submitJobId?: Maybe<Scalars['String']>;
  submitManyJobId?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type Chainlink_Configs_Stddev_Fields = {
  __typename?: 'chainlink_configs_stddev_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Stddev_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** aggregate stddev_pop on columns */
export type Chainlink_Configs_Stddev_Pop_Fields = {
  __typename?: 'chainlink_configs_stddev_pop_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Stddev_Pop_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** aggregate stddev_samp on columns */
export type Chainlink_Configs_Stddev_Samp_Fields = {
  __typename?: 'chainlink_configs_stddev_samp_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Stddev_Samp_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** aggregate sum on columns */
export type Chainlink_Configs_Sum_Fields = {
  __typename?: 'chainlink_configs_sum_fields';
  chainId?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Sum_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** update columns of table "chainlink_configs" */
export enum Chainlink_Configs_Update_Column {
  /** column name */
  ChainId = 'chainId',
  /** column name */
  ConfirmNewAssetJobId = 'confirmNewAssetJobId',
  /** column name */
  Cookie = 'cookie',
  /** column name */
  DebridgeAddr = 'debridgeAddr',
  /** column name */
  EiChainlinkUrl = 'eiChainlinkUrl',
  /** column name */
  EiCiAccesskey = 'eiCiAccesskey',
  /** column name */
  EiCiSecret = 'eiCiSecret',
  /** column name */
  EiIcAccesskey = 'eiIcAccesskey',
  /** column name */
  EiIcSecret = 'eiIcSecret',
  /** column name */
  Network = 'network',
  /** column name */
  Provider = 'provider',
  /** column name */
  SubmitJobId = 'submitJobId',
  /** column name */
  SubmitManyJobId = 'submitManyJobId'
}

/** aggregate var_pop on columns */
export type Chainlink_Configs_Var_Pop_Fields = {
  __typename?: 'chainlink_configs_var_pop_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Var_Pop_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** aggregate var_samp on columns */
export type Chainlink_Configs_Var_Samp_Fields = {
  __typename?: 'chainlink_configs_var_samp_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Var_Samp_Order_By = {
  chainId?: Maybe<Order_By>;
};

/** aggregate variance on columns */
export type Chainlink_Configs_Variance_Fields = {
  __typename?: 'chainlink_configs_variance_fields';
  chainId?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "chainlink_configs" */
export type Chainlink_Configs_Variance_Order_By = {
  chainId?: Maybe<Order_By>;
};

export type Get_Created_Submissions_By_Confirmation_Chain_Id_Args = {
  in_confirmation_chain_id?: Maybe<Scalars['bigint']>;
};

/** mutation root */
export type Mutation_Root = {
  __typename?: 'mutation_root';
  /** delete data from the table: "chainlink_configs" */
  delete_chainlink_configs?: Maybe<Chainlink_Configs_Mutation_Response>;
  /** delete single row from the table: "chainlink_configs" */
  delete_chainlink_configs_by_pk?: Maybe<Chainlink_Configs>;
  /** delete data from the table: "submissions" */
  delete_submissions?: Maybe<Submissions_Mutation_Response>;
  /** delete single row from the table: "submissions" */
  delete_submissions_by_pk?: Maybe<Submissions>;
  /** delete data from the table: "supported_chains" */
  delete_supported_chains?: Maybe<Supported_Chains_Mutation_Response>;
  /** delete single row from the table: "supported_chains" */
  delete_supported_chains_by_pk?: Maybe<Supported_Chains>;
  /** insert data into the table: "chainlink_configs" */
  insert_chainlink_configs?: Maybe<Chainlink_Configs_Mutation_Response>;
  /** insert a single row into the table: "chainlink_configs" */
  insert_chainlink_configs_one?: Maybe<Chainlink_Configs>;
  /** insert data into the table: "submissions" */
  insert_submissions?: Maybe<Submissions_Mutation_Response>;
  /** insert a single row into the table: "submissions" */
  insert_submissions_one?: Maybe<Submissions>;
  /** insert data into the table: "supported_chains" */
  insert_supported_chains?: Maybe<Supported_Chains_Mutation_Response>;
  /** insert a single row into the table: "supported_chains" */
  insert_supported_chains_one?: Maybe<Supported_Chains>;
  /** update data of the table: "chainlink_configs" */
  update_chainlink_configs?: Maybe<Chainlink_Configs_Mutation_Response>;
  /** update single row of the table: "chainlink_configs" */
  update_chainlink_configs_by_pk?: Maybe<Chainlink_Configs>;
  /** update data of the table: "submissions" */
  update_submissions?: Maybe<Submissions_Mutation_Response>;
  /** update single row of the table: "submissions" */
  update_submissions_by_pk?: Maybe<Submissions>;
  /** update data of the table: "supported_chains" */
  update_supported_chains?: Maybe<Supported_Chains_Mutation_Response>;
  /** update single row of the table: "supported_chains" */
  update_supported_chains_by_pk?: Maybe<Supported_Chains>;
};


/** mutation root */
export type Mutation_RootDelete_Chainlink_ConfigsArgs = {
  where: Chainlink_Configs_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Chainlink_Configs_By_PkArgs = {
  chainId: Scalars['bigint'];
};


/** mutation root */
export type Mutation_RootDelete_SubmissionsArgs = {
  where: Submissions_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Submissions_By_PkArgs = {
  id: Scalars['Int'];
};


/** mutation root */
export type Mutation_RootDelete_Supported_ChainsArgs = {
  where: Supported_Chains_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Supported_Chains_By_PkArgs = {
  chainId: Scalars['bigint'];
};


/** mutation root */
export type Mutation_RootInsert_Chainlink_ConfigsArgs = {
  objects: Array<Chainlink_Configs_Insert_Input>;
  on_conflict?: Maybe<Chainlink_Configs_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Chainlink_Configs_OneArgs = {
  object: Chainlink_Configs_Insert_Input;
  on_conflict?: Maybe<Chainlink_Configs_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_SubmissionsArgs = {
  objects: Array<Submissions_Insert_Input>;
  on_conflict?: Maybe<Submissions_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Submissions_OneArgs = {
  object: Submissions_Insert_Input;
  on_conflict?: Maybe<Submissions_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Supported_ChainsArgs = {
  objects: Array<Supported_Chains_Insert_Input>;
  on_conflict?: Maybe<Supported_Chains_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Supported_Chains_OneArgs = {
  object: Supported_Chains_Insert_Input;
  on_conflict?: Maybe<Supported_Chains_On_Conflict>;
};


/** mutation root */
export type Mutation_RootUpdate_Chainlink_ConfigsArgs = {
  _inc?: Maybe<Chainlink_Configs_Inc_Input>;
  _set?: Maybe<Chainlink_Configs_Set_Input>;
  where: Chainlink_Configs_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Chainlink_Configs_By_PkArgs = {
  _inc?: Maybe<Chainlink_Configs_Inc_Input>;
  _set?: Maybe<Chainlink_Configs_Set_Input>;
  pk_columns: Chainlink_Configs_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_SubmissionsArgs = {
  _inc?: Maybe<Submissions_Inc_Input>;
  _set?: Maybe<Submissions_Set_Input>;
  where: Submissions_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Submissions_By_PkArgs = {
  _inc?: Maybe<Submissions_Inc_Input>;
  _set?: Maybe<Submissions_Set_Input>;
  pk_columns: Submissions_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Supported_ChainsArgs = {
  _inc?: Maybe<Supported_Chains_Inc_Input>;
  _set?: Maybe<Supported_Chains_Set_Input>;
  where: Supported_Chains_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Supported_Chains_By_PkArgs = {
  _inc?: Maybe<Supported_Chains_Inc_Input>;
  _set?: Maybe<Supported_Chains_Set_Input>;
  pk_columns: Supported_Chains_Pk_Columns_Input;
};

/** expression to compare columns of type numeric. All fields are combined with logical 'AND'. */
export type Numeric_Comparison_Exp = {
  _eq?: Maybe<Scalars['numeric']>;
  _gt?: Maybe<Scalars['numeric']>;
  _gte?: Maybe<Scalars['numeric']>;
  _in?: Maybe<Array<Scalars['numeric']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['numeric']>;
  _lte?: Maybe<Scalars['numeric']>;
  _neq?: Maybe<Scalars['numeric']>;
  _nin?: Maybe<Array<Scalars['numeric']>>;
};

/** column ordering options */
export enum Order_By {
  /** in the ascending order, nulls last */
  Asc = 'asc',
  /** in the ascending order, nulls first */
  AscNullsFirst = 'asc_nulls_first',
  /** in the ascending order, nulls last */
  AscNullsLast = 'asc_nulls_last',
  /** in the descending order, nulls first */
  Desc = 'desc',
  /** in the descending order, nulls first */
  DescNullsFirst = 'desc_nulls_first',
  /** in the descending order, nulls last */
  DescNullsLast = 'desc_nulls_last'
}

/** query root */
export type Query_Root = {
  __typename?: 'query_root';
  /** fetch data from the table: "chainlink_configs" */
  chainlink_configs: Array<Chainlink_Configs>;
  /** fetch aggregated fields from the table: "chainlink_configs" */
  chainlink_configs_aggregate: Chainlink_Configs_Aggregate;
  /** fetch data from the table: "chainlink_configs" using primary key columns */
  chainlink_configs_by_pk?: Maybe<Chainlink_Configs>;
  /** execute function "get_created_submissions_by_confirmation_chain_id" which returns "submissions" */
  get_created_submissions_by_confirmation_chain_id: Array<Submissions>;
  /** execute function "get_created_submissions_by_confirmation_chain_id" and query aggregates on result of table type "submissions" */
  get_created_submissions_by_confirmation_chain_id_aggregate: Submissions_Aggregate;
  /** fetch data from the table: "submissions" */
  submissions: Array<Submissions>;
  /** fetch aggregated fields from the table: "submissions" */
  submissions_aggregate: Submissions_Aggregate;
  /** fetch data from the table: "submissions" using primary key columns */
  submissions_by_pk?: Maybe<Submissions>;
  /** fetch data from the table: "supported_chains" */
  supported_chains: Array<Supported_Chains>;
  /** fetch aggregated fields from the table: "supported_chains" */
  supported_chains_aggregate: Supported_Chains_Aggregate;
  /** fetch data from the table: "supported_chains" using primary key columns */
  supported_chains_by_pk?: Maybe<Supported_Chains>;
};


/** query root */
export type Query_RootChainlink_ConfigsArgs = {
  distinct_on?: Maybe<Array<Chainlink_Configs_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Chainlink_Configs_Order_By>>;
  where?: Maybe<Chainlink_Configs_Bool_Exp>;
};


/** query root */
export type Query_RootChainlink_Configs_AggregateArgs = {
  distinct_on?: Maybe<Array<Chainlink_Configs_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Chainlink_Configs_Order_By>>;
  where?: Maybe<Chainlink_Configs_Bool_Exp>;
};


/** query root */
export type Query_RootChainlink_Configs_By_PkArgs = {
  chainId: Scalars['bigint'];
};


/** query root */
export type Query_RootGet_Created_Submissions_By_Confirmation_Chain_IdArgs = {
  args: Get_Created_Submissions_By_Confirmation_Chain_Id_Args;
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** query root */
export type Query_RootGet_Created_Submissions_By_Confirmation_Chain_Id_AggregateArgs = {
  args: Get_Created_Submissions_By_Confirmation_Chain_Id_Args;
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** query root */
export type Query_RootSubmissionsArgs = {
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** query root */
export type Query_RootSubmissions_AggregateArgs = {
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** query root */
export type Query_RootSubmissions_By_PkArgs = {
  id: Scalars['Int'];
};


/** query root */
export type Query_RootSupported_ChainsArgs = {
  distinct_on?: Maybe<Array<Supported_Chains_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Supported_Chains_Order_By>>;
  where?: Maybe<Supported_Chains_Bool_Exp>;
};


/** query root */
export type Query_RootSupported_Chains_AggregateArgs = {
  distinct_on?: Maybe<Array<Supported_Chains_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Supported_Chains_Order_By>>;
  where?: Maybe<Supported_Chains_Bool_Exp>;
};


/** query root */
export type Query_RootSupported_Chains_By_PkArgs = {
  chainId: Scalars['bigint'];
};

/** columns and relationships of "submissions" */
export type Submissions = {
  __typename?: 'submissions';
  amount: Scalars['numeric'];
  chainFrom: Scalars['Int'];
  chainTo: Scalars['Int'];
  createdAt: Scalars['timestamptz'];
  debridgeId: Scalars['String'];
  id: Scalars['Int'];
  receiverAddress: Scalars['String'];
  runId: Scalars['String'];
  status: Scalars['Int'];
  submissionId: Scalars['String'];
  txHash: Scalars['String'];
};

/** aggregated selection of "submissions" */
export type Submissions_Aggregate = {
  __typename?: 'submissions_aggregate';
  aggregate?: Maybe<Submissions_Aggregate_Fields>;
  nodes: Array<Submissions>;
};

/** aggregate fields of "submissions" */
export type Submissions_Aggregate_Fields = {
  __typename?: 'submissions_aggregate_fields';
  avg?: Maybe<Submissions_Avg_Fields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<Submissions_Max_Fields>;
  min?: Maybe<Submissions_Min_Fields>;
  stddev?: Maybe<Submissions_Stddev_Fields>;
  stddev_pop?: Maybe<Submissions_Stddev_Pop_Fields>;
  stddev_samp?: Maybe<Submissions_Stddev_Samp_Fields>;
  sum?: Maybe<Submissions_Sum_Fields>;
  var_pop?: Maybe<Submissions_Var_Pop_Fields>;
  var_samp?: Maybe<Submissions_Var_Samp_Fields>;
  variance?: Maybe<Submissions_Variance_Fields>;
};


/** aggregate fields of "submissions" */
export type Submissions_Aggregate_FieldsCountArgs = {
  columns?: Maybe<Array<Submissions_Select_Column>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "submissions" */
export type Submissions_Aggregate_Order_By = {
  avg?: Maybe<Submissions_Avg_Order_By>;
  count?: Maybe<Order_By>;
  max?: Maybe<Submissions_Max_Order_By>;
  min?: Maybe<Submissions_Min_Order_By>;
  stddev?: Maybe<Submissions_Stddev_Order_By>;
  stddev_pop?: Maybe<Submissions_Stddev_Pop_Order_By>;
  stddev_samp?: Maybe<Submissions_Stddev_Samp_Order_By>;
  sum?: Maybe<Submissions_Sum_Order_By>;
  var_pop?: Maybe<Submissions_Var_Pop_Order_By>;
  var_samp?: Maybe<Submissions_Var_Samp_Order_By>;
  variance?: Maybe<Submissions_Variance_Order_By>;
};

/** input type for inserting array relation for remote table "submissions" */
export type Submissions_Arr_Rel_Insert_Input = {
  data: Array<Submissions_Insert_Input>;
  on_conflict?: Maybe<Submissions_On_Conflict>;
};

/** aggregate avg on columns */
export type Submissions_Avg_Fields = {
  __typename?: 'submissions_avg_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "submissions" */
export type Submissions_Avg_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** Boolean expression to filter rows from the table "submissions". All fields are combined with a logical 'AND'. */
export type Submissions_Bool_Exp = {
  _and?: Maybe<Array<Maybe<Submissions_Bool_Exp>>>;
  _not?: Maybe<Submissions_Bool_Exp>;
  _or?: Maybe<Array<Maybe<Submissions_Bool_Exp>>>;
  amount?: Maybe<Numeric_Comparison_Exp>;
  chainFrom?: Maybe<Int_Comparison_Exp>;
  chainTo?: Maybe<Int_Comparison_Exp>;
  createdAt?: Maybe<Timestamptz_Comparison_Exp>;
  debridgeId?: Maybe<String_Comparison_Exp>;
  id?: Maybe<Int_Comparison_Exp>;
  receiverAddress?: Maybe<String_Comparison_Exp>;
  runId?: Maybe<String_Comparison_Exp>;
  status?: Maybe<Int_Comparison_Exp>;
  submissionId?: Maybe<String_Comparison_Exp>;
  txHash?: Maybe<String_Comparison_Exp>;
};

/** unique or primary key constraints on table "submissions" */
export enum Submissions_Constraint {
  /** unique or primary key constraint */
  SubmissionsPkey = 'submissions_pkey',
  /** unique or primary key constraint */
  SubmissionsSubmissionIdKey = 'submissions_submission_id_key'
}

/** input type for incrementing integer column in table "submissions" */
export type Submissions_Inc_Input = {
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  id?: Maybe<Scalars['Int']>;
  status?: Maybe<Scalars['Int']>;
};

/** input type for inserting data into table "submissions" */
export type Submissions_Insert_Input = {
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  debridgeId?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  receiverAddress?: Maybe<Scalars['String']>;
  runId?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['Int']>;
  submissionId?: Maybe<Scalars['String']>;
  txHash?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type Submissions_Max_Fields = {
  __typename?: 'submissions_max_fields';
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  debridgeId?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  receiverAddress?: Maybe<Scalars['String']>;
  runId?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['Int']>;
  submissionId?: Maybe<Scalars['String']>;
  txHash?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "submissions" */
export type Submissions_Max_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  createdAt?: Maybe<Order_By>;
  debridgeId?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  receiverAddress?: Maybe<Order_By>;
  runId?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
  submissionId?: Maybe<Order_By>;
  txHash?: Maybe<Order_By>;
};

/** aggregate min on columns */
export type Submissions_Min_Fields = {
  __typename?: 'submissions_min_fields';
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  debridgeId?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  receiverAddress?: Maybe<Scalars['String']>;
  runId?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['Int']>;
  submissionId?: Maybe<Scalars['String']>;
  txHash?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "submissions" */
export type Submissions_Min_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  createdAt?: Maybe<Order_By>;
  debridgeId?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  receiverAddress?: Maybe<Order_By>;
  runId?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
  submissionId?: Maybe<Order_By>;
  txHash?: Maybe<Order_By>;
};

/** response of any mutation on the table "submissions" */
export type Submissions_Mutation_Response = {
  __typename?: 'submissions_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Submissions>;
};

/** input type for inserting object relation for remote table "submissions" */
export type Submissions_Obj_Rel_Insert_Input = {
  data: Submissions_Insert_Input;
  on_conflict?: Maybe<Submissions_On_Conflict>;
};

/** on conflict condition type for table "submissions" */
export type Submissions_On_Conflict = {
  constraint: Submissions_Constraint;
  update_columns: Array<Submissions_Update_Column>;
  where?: Maybe<Submissions_Bool_Exp>;
};

/** ordering options when selecting data from "submissions" */
export type Submissions_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  createdAt?: Maybe<Order_By>;
  debridgeId?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  receiverAddress?: Maybe<Order_By>;
  runId?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
  submissionId?: Maybe<Order_By>;
  txHash?: Maybe<Order_By>;
};

/** primary key columns input for table: "submissions" */
export type Submissions_Pk_Columns_Input = {
  id: Scalars['Int'];
};

/** select columns of table "submissions" */
export enum Submissions_Select_Column {
  /** column name */
  Amount = 'amount',
  /** column name */
  ChainFrom = 'chainFrom',
  /** column name */
  ChainTo = 'chainTo',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  DebridgeId = 'debridgeId',
  /** column name */
  Id = 'id',
  /** column name */
  ReceiverAddress = 'receiverAddress',
  /** column name */
  RunId = 'runId',
  /** column name */
  Status = 'status',
  /** column name */
  SubmissionId = 'submissionId',
  /** column name */
  TxHash = 'txHash'
}

/** input type for updating data in table "submissions" */
export type Submissions_Set_Input = {
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  debridgeId?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  receiverAddress?: Maybe<Scalars['String']>;
  runId?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['Int']>;
  submissionId?: Maybe<Scalars['String']>;
  txHash?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type Submissions_Stddev_Fields = {
  __typename?: 'submissions_stddev_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "submissions" */
export type Submissions_Stddev_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** aggregate stddev_pop on columns */
export type Submissions_Stddev_Pop_Fields = {
  __typename?: 'submissions_stddev_pop_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "submissions" */
export type Submissions_Stddev_Pop_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** aggregate stddev_samp on columns */
export type Submissions_Stddev_Samp_Fields = {
  __typename?: 'submissions_stddev_samp_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "submissions" */
export type Submissions_Stddev_Samp_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** aggregate sum on columns */
export type Submissions_Sum_Fields = {
  __typename?: 'submissions_sum_fields';
  amount?: Maybe<Scalars['numeric']>;
  chainFrom?: Maybe<Scalars['Int']>;
  chainTo?: Maybe<Scalars['Int']>;
  id?: Maybe<Scalars['Int']>;
  status?: Maybe<Scalars['Int']>;
};

/** order by sum() on columns of table "submissions" */
export type Submissions_Sum_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** update columns of table "submissions" */
export enum Submissions_Update_Column {
  /** column name */
  Amount = 'amount',
  /** column name */
  ChainFrom = 'chainFrom',
  /** column name */
  ChainTo = 'chainTo',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  DebridgeId = 'debridgeId',
  /** column name */
  Id = 'id',
  /** column name */
  ReceiverAddress = 'receiverAddress',
  /** column name */
  RunId = 'runId',
  /** column name */
  Status = 'status',
  /** column name */
  SubmissionId = 'submissionId',
  /** column name */
  TxHash = 'txHash'
}

/** aggregate var_pop on columns */
export type Submissions_Var_Pop_Fields = {
  __typename?: 'submissions_var_pop_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "submissions" */
export type Submissions_Var_Pop_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** aggregate var_samp on columns */
export type Submissions_Var_Samp_Fields = {
  __typename?: 'submissions_var_samp_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "submissions" */
export type Submissions_Var_Samp_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** aggregate variance on columns */
export type Submissions_Variance_Fields = {
  __typename?: 'submissions_variance_fields';
  amount?: Maybe<Scalars['Float']>;
  chainFrom?: Maybe<Scalars['Float']>;
  chainTo?: Maybe<Scalars['Float']>;
  id?: Maybe<Scalars['Float']>;
  status?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "submissions" */
export type Submissions_Variance_Order_By = {
  amount?: Maybe<Order_By>;
  chainFrom?: Maybe<Order_By>;
  chainTo?: Maybe<Order_By>;
  id?: Maybe<Order_By>;
  status?: Maybe<Order_By>;
};

/** subscription root */
export type Subscription_Root = {
  __typename?: 'subscription_root';
  /** fetch data from the table: "chainlink_configs" */
  chainlink_configs: Array<Chainlink_Configs>;
  /** fetch aggregated fields from the table: "chainlink_configs" */
  chainlink_configs_aggregate: Chainlink_Configs_Aggregate;
  /** fetch data from the table: "chainlink_configs" using primary key columns */
  chainlink_configs_by_pk?: Maybe<Chainlink_Configs>;
  /** execute function "get_created_submissions_by_confirmation_chain_id" which returns "submissions" */
  get_created_submissions_by_confirmation_chain_id: Array<Submissions>;
  /** execute function "get_created_submissions_by_confirmation_chain_id" and query aggregates on result of table type "submissions" */
  get_created_submissions_by_confirmation_chain_id_aggregate: Submissions_Aggregate;
  /** fetch data from the table: "submissions" */
  submissions: Array<Submissions>;
  /** fetch aggregated fields from the table: "submissions" */
  submissions_aggregate: Submissions_Aggregate;
  /** fetch data from the table: "submissions" using primary key columns */
  submissions_by_pk?: Maybe<Submissions>;
  /** fetch data from the table: "supported_chains" */
  supported_chains: Array<Supported_Chains>;
  /** fetch aggregated fields from the table: "supported_chains" */
  supported_chains_aggregate: Supported_Chains_Aggregate;
  /** fetch data from the table: "supported_chains" using primary key columns */
  supported_chains_by_pk?: Maybe<Supported_Chains>;
};


/** subscription root */
export type Subscription_RootChainlink_ConfigsArgs = {
  distinct_on?: Maybe<Array<Chainlink_Configs_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Chainlink_Configs_Order_By>>;
  where?: Maybe<Chainlink_Configs_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootChainlink_Configs_AggregateArgs = {
  distinct_on?: Maybe<Array<Chainlink_Configs_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Chainlink_Configs_Order_By>>;
  where?: Maybe<Chainlink_Configs_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootChainlink_Configs_By_PkArgs = {
  chainId: Scalars['bigint'];
};


/** subscription root */
export type Subscription_RootGet_Created_Submissions_By_Confirmation_Chain_IdArgs = {
  args: Get_Created_Submissions_By_Confirmation_Chain_Id_Args;
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootGet_Created_Submissions_By_Confirmation_Chain_Id_AggregateArgs = {
  args: Get_Created_Submissions_By_Confirmation_Chain_Id_Args;
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootSubmissionsArgs = {
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootSubmissions_AggregateArgs = {
  distinct_on?: Maybe<Array<Submissions_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Submissions_Order_By>>;
  where?: Maybe<Submissions_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootSubmissions_By_PkArgs = {
  id: Scalars['Int'];
};


/** subscription root */
export type Subscription_RootSupported_ChainsArgs = {
  distinct_on?: Maybe<Array<Supported_Chains_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Supported_Chains_Order_By>>;
  where?: Maybe<Supported_Chains_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootSupported_Chains_AggregateArgs = {
  distinct_on?: Maybe<Array<Supported_Chains_Select_Column>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<Supported_Chains_Order_By>>;
  where?: Maybe<Supported_Chains_Bool_Exp>;
};


/** subscription root */
export type Subscription_RootSupported_Chains_By_PkArgs = {
  chainId: Scalars['bigint'];
};

/** columns and relationships of "supported_chains" */
export type Supported_Chains = {
  __typename?: 'supported_chains';
  chainId: Scalars['bigint'];
  /** An object relationship */
  config?: Maybe<Chainlink_Configs>;
  /** Номер блокчейна в сети (eth, etc), куда мы делаем перевод. Куда мы агрегируем подтверждения. */
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock: Scalars['Int'];
  network: Scalars['String'];
};

/** aggregated selection of "supported_chains" */
export type Supported_Chains_Aggregate = {
  __typename?: 'supported_chains_aggregate';
  aggregate?: Maybe<Supported_Chains_Aggregate_Fields>;
  nodes: Array<Supported_Chains>;
};

/** aggregate fields of "supported_chains" */
export type Supported_Chains_Aggregate_Fields = {
  __typename?: 'supported_chains_aggregate_fields';
  avg?: Maybe<Supported_Chains_Avg_Fields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<Supported_Chains_Max_Fields>;
  min?: Maybe<Supported_Chains_Min_Fields>;
  stddev?: Maybe<Supported_Chains_Stddev_Fields>;
  stddev_pop?: Maybe<Supported_Chains_Stddev_Pop_Fields>;
  stddev_samp?: Maybe<Supported_Chains_Stddev_Samp_Fields>;
  sum?: Maybe<Supported_Chains_Sum_Fields>;
  var_pop?: Maybe<Supported_Chains_Var_Pop_Fields>;
  var_samp?: Maybe<Supported_Chains_Var_Samp_Fields>;
  variance?: Maybe<Supported_Chains_Variance_Fields>;
};


/** aggregate fields of "supported_chains" */
export type Supported_Chains_Aggregate_FieldsCountArgs = {
  columns?: Maybe<Array<Supported_Chains_Select_Column>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "supported_chains" */
export type Supported_Chains_Aggregate_Order_By = {
  avg?: Maybe<Supported_Chains_Avg_Order_By>;
  count?: Maybe<Order_By>;
  max?: Maybe<Supported_Chains_Max_Order_By>;
  min?: Maybe<Supported_Chains_Min_Order_By>;
  stddev?: Maybe<Supported_Chains_Stddev_Order_By>;
  stddev_pop?: Maybe<Supported_Chains_Stddev_Pop_Order_By>;
  stddev_samp?: Maybe<Supported_Chains_Stddev_Samp_Order_By>;
  sum?: Maybe<Supported_Chains_Sum_Order_By>;
  var_pop?: Maybe<Supported_Chains_Var_Pop_Order_By>;
  var_samp?: Maybe<Supported_Chains_Var_Samp_Order_By>;
  variance?: Maybe<Supported_Chains_Variance_Order_By>;
};

/** input type for inserting array relation for remote table "supported_chains" */
export type Supported_Chains_Arr_Rel_Insert_Input = {
  data: Array<Supported_Chains_Insert_Input>;
  on_conflict?: Maybe<Supported_Chains_On_Conflict>;
};

/** aggregate avg on columns */
export type Supported_Chains_Avg_Fields = {
  __typename?: 'supported_chains_avg_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "supported_chains" */
export type Supported_Chains_Avg_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** Boolean expression to filter rows from the table "supported_chains". All fields are combined with a logical 'AND'. */
export type Supported_Chains_Bool_Exp = {
  _and?: Maybe<Array<Maybe<Supported_Chains_Bool_Exp>>>;
  _not?: Maybe<Supported_Chains_Bool_Exp>;
  _or?: Maybe<Array<Maybe<Supported_Chains_Bool_Exp>>>;
  chainId?: Maybe<Bigint_Comparison_Exp>;
  config?: Maybe<Chainlink_Configs_Bool_Exp>;
  confirmationChainId?: Maybe<Bigint_Comparison_Exp>;
  latestBlock?: Maybe<Int_Comparison_Exp>;
  network?: Maybe<String_Comparison_Exp>;
};

/** unique or primary key constraints on table "supported_chains" */
export enum Supported_Chains_Constraint {
  /** unique or primary key constraint */
  SupportedChainsPkey = 'supported_chains_pkey'
}

/** input type for incrementing integer column in table "supported_chains" */
export type Supported_Chains_Inc_Input = {
  chainId?: Maybe<Scalars['bigint']>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
};

/** input type for inserting data into table "supported_chains" */
export type Supported_Chains_Insert_Input = {
  chainId?: Maybe<Scalars['bigint']>;
  config?: Maybe<Chainlink_Configs_Obj_Rel_Insert_Input>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
  network?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type Supported_Chains_Max_Fields = {
  __typename?: 'supported_chains_max_fields';
  chainId?: Maybe<Scalars['bigint']>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
  network?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "supported_chains" */
export type Supported_Chains_Max_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
};

/** aggregate min on columns */
export type Supported_Chains_Min_Fields = {
  __typename?: 'supported_chains_min_fields';
  chainId?: Maybe<Scalars['bigint']>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
  network?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "supported_chains" */
export type Supported_Chains_Min_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
};

/** response of any mutation on the table "supported_chains" */
export type Supported_Chains_Mutation_Response = {
  __typename?: 'supported_chains_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Supported_Chains>;
};

/** input type for inserting object relation for remote table "supported_chains" */
export type Supported_Chains_Obj_Rel_Insert_Input = {
  data: Supported_Chains_Insert_Input;
  on_conflict?: Maybe<Supported_Chains_On_Conflict>;
};

/** on conflict condition type for table "supported_chains" */
export type Supported_Chains_On_Conflict = {
  constraint: Supported_Chains_Constraint;
  update_columns: Array<Supported_Chains_Update_Column>;
  where?: Maybe<Supported_Chains_Bool_Exp>;
};

/** ordering options when selecting data from "supported_chains" */
export type Supported_Chains_Order_By = {
  chainId?: Maybe<Order_By>;
  config?: Maybe<Chainlink_Configs_Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
  network?: Maybe<Order_By>;
};

/** primary key columns input for table: "supported_chains" */
export type Supported_Chains_Pk_Columns_Input = {
  chainId: Scalars['bigint'];
};

/** select columns of table "supported_chains" */
export enum Supported_Chains_Select_Column {
  /** column name */
  ChainId = 'chainId',
  /** column name */
  ConfirmationChainId = 'confirmationChainId',
  /** column name */
  LatestBlock = 'latestBlock',
  /** column name */
  Network = 'network'
}

/** input type for updating data in table "supported_chains" */
export type Supported_Chains_Set_Input = {
  chainId?: Maybe<Scalars['bigint']>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
  network?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type Supported_Chains_Stddev_Fields = {
  __typename?: 'supported_chains_stddev_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "supported_chains" */
export type Supported_Chains_Stddev_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** aggregate stddev_pop on columns */
export type Supported_Chains_Stddev_Pop_Fields = {
  __typename?: 'supported_chains_stddev_pop_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "supported_chains" */
export type Supported_Chains_Stddev_Pop_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** aggregate stddev_samp on columns */
export type Supported_Chains_Stddev_Samp_Fields = {
  __typename?: 'supported_chains_stddev_samp_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "supported_chains" */
export type Supported_Chains_Stddev_Samp_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** aggregate sum on columns */
export type Supported_Chains_Sum_Fields = {
  __typename?: 'supported_chains_sum_fields';
  chainId?: Maybe<Scalars['bigint']>;
  confirmationChainId?: Maybe<Scalars['bigint']>;
  latestBlock?: Maybe<Scalars['Int']>;
};

/** order by sum() on columns of table "supported_chains" */
export type Supported_Chains_Sum_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** update columns of table "supported_chains" */
export enum Supported_Chains_Update_Column {
  /** column name */
  ChainId = 'chainId',
  /** column name */
  ConfirmationChainId = 'confirmationChainId',
  /** column name */
  LatestBlock = 'latestBlock',
  /** column name */
  Network = 'network'
}

/** aggregate var_pop on columns */
export type Supported_Chains_Var_Pop_Fields = {
  __typename?: 'supported_chains_var_pop_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "supported_chains" */
export type Supported_Chains_Var_Pop_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** aggregate var_samp on columns */
export type Supported_Chains_Var_Samp_Fields = {
  __typename?: 'supported_chains_var_samp_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "supported_chains" */
export type Supported_Chains_Var_Samp_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** aggregate variance on columns */
export type Supported_Chains_Variance_Fields = {
  __typename?: 'supported_chains_variance_fields';
  chainId?: Maybe<Scalars['Float']>;
  confirmationChainId?: Maybe<Scalars['Float']>;
  latestBlock?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "supported_chains" */
export type Supported_Chains_Variance_Order_By = {
  chainId?: Maybe<Order_By>;
  confirmationChainId?: Maybe<Order_By>;
  latestBlock?: Maybe<Order_By>;
};

/** expression to compare columns of type timestamptz. All fields are combined with logical 'AND'. */
export type Timestamptz_Comparison_Exp = {
  _eq?: Maybe<Scalars['timestamptz']>;
  _gt?: Maybe<Scalars['timestamptz']>;
  _gte?: Maybe<Scalars['timestamptz']>;
  _in?: Maybe<Array<Scalars['timestamptz']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['timestamptz']>;
  _lte?: Maybe<Scalars['timestamptz']>;
  _neq?: Maybe<Scalars['timestamptz']>;
  _nin?: Maybe<Array<Scalars['timestamptz']>>;
};

export type GetChainDetailsQueryVariables = Exact<{
  chainId: Scalars['bigint'];
}>;


export type GetChainDetailsQuery = { __typename?: 'query_root', supported_chains_by_pk?: Maybe<{ __typename?: 'supported_chains', chainId: any, latestBlock: number, config?: Maybe<{ __typename?: 'chainlink_configs', eiChainlinkUrl: string, eiCiAccesskey: string, eiCiSecret: string, eiIcAccesskey: string, eiIcSecret: string, provider: string, debridgeAddr: string, network: string }> }> };

export type GetChainSubmissionDetailsQueryVariables = Exact<{
  confirmationChainId: Scalars['bigint'];
  submissionId: Scalars['Int'];
}>;


export type GetChainSubmissionDetailsQuery = { __typename?: 'query_root', supported_chains: Array<{ __typename?: 'supported_chains', config?: Maybe<{ __typename?: 'chainlink_configs', chainId: any, eiChainlinkUrl: string, eiCiAccesskey: string, eiCiSecret: string, eiIcAccesskey: string, eiIcSecret: string }> }>, submissions_by_pk?: Maybe<{ __typename?: 'submissions', id: number }> };

export type GetChainsConfigQueryVariables = Exact<{ [key: string]: never; }>;


export type GetChainsConfigQuery = { __typename?: 'query_root', supported_chains: Array<{ __typename?: 'supported_chains', chainId: any, network: string }>, chainlink_configs: Array<{ __typename?: 'chainlink_configs', chainId: any, network: string }> };

export type GetNewSubmissionsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetNewSubmissionsQuery = { __typename?: 'query_root', submissions: Array<{ __typename?: 'submissions', submissionId: string, chainTo: number }>, chainlink_configs: Array<{ __typename?: 'chainlink_configs', chainId: any, submitJobId: string, submitManyJobId: string, confirmNewAssetJobId: string, eiChainlinkUrl: string, eiIcAccesskey: string, eiIcSecret: string }> };

export type GetCreatedSubmissionsByConfiramationChainIdQueryVariables = Exact<{
  confirmationChainId: Scalars['bigint'];
}>;


export type GetCreatedSubmissionsByConfiramationChainIdQuery = { __typename?: 'query_root', get_created_submissions_by_confirmation_chain_id: Array<{ __typename?: 'submissions', submissionId: string }> };

export type InsertSubmissionMutationVariables = Exact<{
  object: Submissions_Insert_Input;
}>;


export type InsertSubmissionMutation = { __typename?: 'mutation_root', insert_submissions_one?: Maybe<{ __typename?: 'submissions', id: number }> };

export type UpdateChainLatestBlockMutationVariables = Exact<{
  chainId: Scalars['bigint'];
  latestBlock: Scalars['Int'];
}>;


export type UpdateChainLatestBlockMutation = { __typename?: 'mutation_root', update_supported_chains_by_pk?: Maybe<{ __typename?: 'supported_chains', latestBlock: number }> };

export type UpdateSubmissionsMutationVariables = Exact<{
  runId: Scalars['String'];
  status: Scalars['Int'];
  submissionIds: Array<Scalars['String']> | Scalars['String'];
}>;


export type UpdateSubmissionsMutation = { __typename?: 'mutation_root', update_submissions?: Maybe<{ __typename?: 'submissions_mutation_response', affected_rows: number }> };
