# Binary Tree Training System Design

## Goal

Build the binary tree topic into the first long-term algorithm training system in this repository.

The system should serve four purposes at once:

- Interview readiness: high-frequency problems, clear explanations, and concise recall.
- Algorithm understanding: traversal choices, recursion invariants, return-value design, and proof of coverage.
- Practice: focused exercises with tests and review paths.
- Visualization: interactive explanations for problems where state transitions are hard to see from code alone.

This is not a full binary tree encyclopedia. The scope is limited to common, high-frequency, high-training-value problems.

## Scope

The first version should cover representative binary tree problems from `interview/` and `doc/algorithm-visualizations/`.

Priority groups:

- Basic traversal: preorder, inorder, postorder, level order.
- Depth and path problems: max depth, path sum, path collection, path aggregation.
- Structure mutation: invert tree, flatten tree, merge two binary trees.
- BST reasoning: validate BST, kth smallest, BST iterator.
- Construction and encoding: build tree, serialize and deserialize tree.
- Lowest common ancestor: standard LCA and variants.
- Return-value recursion: diameter, max path sum, subtree checks.

Problems outside these groups can stay in the repository, but they should not drive this system unless they are common in interviews or teach a reusable binary tree pattern.

## Deliverables

1. Binary tree gap inventory

   Create a compact inventory of common binary tree problems with these columns:

   - Problem number and title
   - Training category
   - Go solution status
   - Test status
   - Markdown note status
   - Four-question structure status
   - Visualization status
   - Priority
   - Next action

2. Representative problem sample

   Use `617.merge-two-binary-trees` as the first complete sample because it is already modified in the working tree and has clear training value:

   - Synchronous recursion over two trees
   - Empty-node branching
   - In-place mutation versus new-node construction
   - Recursion termination and coverage proof

3. Binary tree training index

   Add or update a binary tree topic document that organizes common problems by skill, not by problem number. The document should help answer:

   - What pattern does this problem train?
   - Which traversal or recursion shape should I consider?
   - What should the recursive call return?
   - What are the common mistakes?
   - Which problem should I practice next?

4. Problem note standard

   Common binary tree notes should converge on this structure:

   - Recognition signals
   - Core invariant
   - Decision process
   - Why the solution does not miss answers
   - Interview explanation
   - Common mistakes
   - Complexity
   - Related transfer problems

5. Visualization policy

   Do not visualize every binary tree problem. Prioritize visualizations when they clarify recursion state, subtree returns, path accumulation, tree mutation, or serialization state.

## Execution Order

1. Inventory common binary tree problems and classify gaps.
2. Complete `617.merge-two-binary-trees` as the first representative sample.
3. Create or update the binary tree training index from the inventory.
4. Upgrade the highest-value problem notes and tests in small batches.
5. Add visualizations only for problems where diagrams materially improve understanding.

Each concrete problem change should follow the repository rule: one problem, one commit.

## Non-Goals

- Do not cover every tree-related LeetCode problem.
- Do not rewrite unrelated algorithm templates.
- Do not create visualizations just to increase count.
- Do not mix many problem upgrades into one commit.
- Do not optimize for file count over learning value.

## Success Criteria

The binary tree topic is successful when a learner can:

- Pick a common binary tree problem and know which pattern it belongs to.
- Explain why a traversal or recursive return value is correct.
- Run focused tests for the implementation.
- Review the problem through a consistent note format.
- Use visualizations for the few problems where code alone is not enough.
- Move from one problem to the next through a deliberate training path.

## First Implementation Plan Boundary

The first implementation plan should stop at:

- Producing the binary tree gap inventory.
- Finishing the `617.merge-two-binary-trees` sample.
- Drafting the first version of the binary tree training index.

Later batches can upgrade additional notes, tests, and visualizations based on the inventory.
