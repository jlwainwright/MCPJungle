# Claude Code Agentic Loop Workflow

## Overview
This workflow implements a 5-iteration planning-coding cycle where each completed section is reviewed by a planning agent, enhanced, then passed to a coding agent. The process ensures continuous improvement and documentation.

## Workflow Structure

### 1. Initialization
When starting a new task with this workflow:
1. Copy this file to your project root
2. Create a `.claude_workflow/` directory
3. Run: `claude "Initialize agentic workflow for: [YOUR TASK DESCRIPTION]"`

### 2. Iteration Cycle (Repeat 5 times)

#### Phase A: Planning Review
```
Role: Planning Agent
Input: Previous iteration's completed work (or initial task for iteration 1)
Output: Structured feedback and enhancement recommendations

Prompt Template:
"As a planning agent, review the completed work for iteration [N]:
- Analyze what was accomplished
- Identify areas for improvement
- Suggest specific enhancements
- Define success criteria for next coding phase
- Provide structured JSON feedback"
```

#### Phase B: Coding Implementation
```
Role: Coding Agent
Input: Planning agent's feedback and recommendations
Output: Implemented enhancements with documentation

Prompt Template:
"As a coding agent for iteration [N], implement the planning feedback:
- Follow all enhancement recommendations
- Make incremental, testable changes
- Document all modifications
- Ensure backward compatibility
- Summarize accomplishments"
```

### 3. Workflow Rules

#### For Planning Agent:
1. Always output structured JSON feedback
2. Be specific about improvements needed
3. Consider technical debt and maintainability
4. Validate against original task requirements
5. Document reasoning for recommendations

#### For Coding Agent:
1. Implement ONLY what planning agent specified
2. Create tests for new functionality
3. Update documentation as you code
4. Commit changes with descriptive messages
5. Report completion status clearly

### 4. Context Management

#### Iteration Context Structure:
```markdown
## Iteration [N] Context
### Previous Work
- [Summary of completed features]
- [Files modified]
- [Tests added/passed]

### Planning Feedback
- [JSON feedback from planning agent]

### Current Goals
1. [Specific goal 1]
2. [Specific goal 2]
...

### Constraints
- [Time/scope constraints]
- [Technical constraints]
```

### 5. Documentation Requirements

#### Per Iteration:
```markdown
## Iteration [N] Results
**Duration**: [time taken]
**Planning Focus**: [main improvements targeted]
**Implementation Summary**: 
- ✓ [Completed item 1]
- ✓ [Completed item 2]
- ⚠ [Partial/blocked item]
**Key Decisions**: [Technical choices made]
**Challenges**: [Issues encountered]
**Next Steps**: [From planning agent]
```

#### Final Summary:
```markdown
# Workflow Completion Report

## Task Overview
- **Original Task**: [description]
- **Total Iterations**: 5
- **Total Duration**: [time]
- **Success Rate**: [percentage]

## Key Achievements
1. [Major accomplishment 1]
2. [Major accomplishment 2]
...

## Technical Decisions Log
| Decision | Reasoning | Impact |
|----------|-----------|---------|
| [Choice] | [Why] | [Result] |

## Lessons Learned
- [Insight 1]
- [Insight 2]

## Recommendations
- [Future improvement 1]
- [Future improvement 2]
```

### 6. Practical Usage

#### Starting a Workflow:
```bash
# 1. Initialize in Claude Code
claude "Start agentic workflow for: Implement user authentication system"

# 2. First planning phase
claude "As planning agent, analyze the task and provide initial implementation plan as JSON"

# 3. First coding phase  
claude "As coding agent, implement iteration 1 based on planning feedback"

# 4. Continue cycle...
```

#### Resuming a Workflow:
```bash
# Continue from last checkpoint
claude -c

# Or explicitly state position
claude "Resume agentic workflow at iteration 3, planning phase"
```

### 7. Checkpoint System

After each iteration, create a checkpoint:
```bash
# Git checkpoint
git add -A
git commit -m "Checkpoint: Iteration [N] - [summary]"
git tag iteration-[N]

# Document checkpoint
echo "Iteration [N] completed at $(date)" >> .claude_workflow/checkpoints.log
```

### 8. Emergency Procedures

#### If workflow breaks:
1. Check `.claude_workflow/checkpoints.log`
2. Review last git commit
3. Resume with: `claude "Resume agentic workflow from last checkpoint"`

#### If planning/coding agents conflict:
1. Prioritize planning agent's latest feedback
2. Document conflict in results
3. Adjust approach in next iteration

### 9. Workflow Configuration

```yaml
# .claude_workflow/config.yaml
workflow:
  iterations: 5
  task: "Your task description"
  started: "timestamp"
  
planning_agent:
  output_format: "json"
  review_depth: "comprehensive"
  focus_areas:
    - "functionality"
    - "performance"
    - "maintainability"
    - "security"
    
coding_agent:
  commit_style: "conventional"
  test_requirement: "mandatory"
  documentation: "inline"
  
checkpoints:
  auto_commit: true
  tag_iterations: true
```

### 10. Success Metrics

Track these metrics throughout the workflow:
- Lines of code added/removed
- Test coverage percentage
- Documentation completeness
- Performance benchmarks
- Technical debt introduced/resolved

## Quick Reference Commands

```bash
# Start workflow
claude "Initialize agentic workflow for: [TASK]"

# Planning phase
claude "As planning agent for iteration [N], review and provide JSON feedback"

# Coding phase
claude "As coding agent for iteration [N], implement planning feedback"

# Check status
claude "Show current agentic workflow status"

# Generate final report
claude "Generate final workflow report with all iterations"
```

## Important Notes

1. **Maintain Role Clarity**: Always specify if you're acting as planning or coding agent
2. **Document Everything**: Every decision and change should be logged
3. **Incremental Progress**: Each iteration should produce working code
4. **Continuous Integration**: Run tests after each coding phase
5. **Context Preservation**: Use CLAUDE.md imports to maintain context

This workflow ensures systematic progress with continuous improvement and comprehensive documentation throughout the development process.