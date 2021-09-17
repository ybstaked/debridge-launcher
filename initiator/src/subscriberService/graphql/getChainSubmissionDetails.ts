import { gql } from "graphql-request";



export const GET_CHAIN_SUBMISSION_DETAILS_QUERY = gql`
query getChainSubmissionDetails($confirmationChainId: bigint!, $submissionId: Int!) {
  supported_chains(where: {confirmationChainId: {_eq: $confirmationChainId}}) {
    config {
      chainId
      eiChainlinkUrl
      eiCiAccesskey
      eiCiSecret
      eiIcAccesskey
      eiIcSecret
    }
  }

  submissions_by_pk(id: $submissionId) {
    id
  }
}
`;
