# AIW3 NFT Team Conventions

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Establishes team collaboration workflows, communication protocols, and development processes.

---

## Table of Contents

1. [Git Workflow & Branching Strategy](#git-workflow--branching-strategy)
2. [Code Review Process](#code-review-process)
3. [Communication Protocols](#communication-protocols)
4. [Meeting & Planning Cadences](#meeting--planning-cadences)
5. [Documentation Standards](#documentation-standards)
6. [Issue Tracking & Project Management](#issue-tracking--project-management)
7. [Release & Deployment Coordination](#release--deployment-coordination)
8. [Knowledge Sharing Practices](#knowledge-sharing-practices)
9. [Quality Assurance Coordination](#quality-assurance-coordination)
10. [Emergency Response Procedures](#emergency-response-procedures)

---

## Git Workflow & Branching Strategy

### Branch Naming Conventions

```
main                    # Production-ready code
develop                 # Integration branch for features
feature/NFT-123-claim   # Feature branches (Jira ticket + description)
bugfix/NFT-456-fix      # Bug fix branches
hotfix/NFT-789-urgent   # Critical production fixes
release/v1.2.0          # Release preparation branches
```

### Branching Workflow

#### 1. Feature Development
```bash
# Start from develop branch
git checkout develop
git pull origin develop

# Create feature branch
git checkout -b feature/NFT-123-nft-claiming-api

# Work on feature, commit regularly
git add .
git commit -m "feat(nft): implement NFT claiming validation logic

- Add qualification check for trading volume
- Implement Redis caching for user qualifications
- Add comprehensive input validation

Closes NFT-123"

# Push and create PR
git push origin feature/NFT-123-nft-claiming-api
```

#### 2. Commit Message Standards
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
**Scopes**: `nft`, `web3`, `api`, `frontend`, `db`, `deploy`, `test`

**Examples**:
```
feat(nft): add NFT upgrade burn-and-mint workflow

- Implement burn transaction for old NFT
- Add metadata upload to IPFS for new NFT
- Integrate with Web3Service for minting
- Add comprehensive error handling and rollback

Closes NFT-234

fix(web3): resolve Solana RPC timeout issues

- Increase connection timeout to 30 seconds
- Add exponential backoff retry logic
- Implement circuit breaker for RPC failures

Fixes NFT-456

docs(api): update NFT claiming endpoint documentation

- Add request/response examples
- Document error codes and scenarios
- Update authentication requirements

Updates NFT-567
```

### Pull Request Process

#### PR Template
```markdown
## Description
Brief description of changes and motivation.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed
- [ ] Performance impact assessed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Code is commented, particularly in hard-to-understand areas
- [ ] Documentation updated
- [ ] No new warnings introduced

## Related Issues
Closes #123
Relates to #456
```

#### Review Requirements
- **Minimum 2 approvals** for production code
- **1 approval** for documentation/test changes
- **All CI checks must pass** before merge
- **Conflicts must be resolved** by PR author

---

## Code Review Process

### Review Guidelines

#### For Reviewers
**Review Checklist**:
- [ ] **Functionality**: Does the code do what it's supposed to do?
- [ ] **Architecture**: Does it follow SOLID principles and system design?
- [ ] **Security**: Are there any security vulnerabilities?
- [ ] **Performance**: Are there any performance implications?
- [ ] **Testing**: Are there adequate tests?
- [ ] **Documentation**: Is the code well-documented?
- [ ] **Standards**: Does it follow our coding standards?

**Review Response Time**:
- **Critical/Hotfix**: 2 hours
- **Feature**: 24 hours
- **Documentation**: 48 hours

#### For PR Authors
**Before Requesting Review**:
- [ ] Self-review completed
- [ ] All tests pass locally
- [ ] Documentation updated
- [ ] Commit messages follow conventions
- [ ] PR description is complete

**Responding to Feedback**:
- **Address all comments** before requesting re-review
- **Explain decisions** when disagreeing with feedback
- **Update PR description** if scope changes
- **Resolve conversations** after addressing feedback

### Review Types

#### 1. Architecture Review (Senior Developers)
- System design alignment
- Service integration patterns
- Database schema changes
- Performance implications

#### 2. Security Review (Security Lead)
- Authentication/authorization
- Input validation
- Sensitive data handling
- External API integrations

#### 3. Code Quality Review (All Developers)
- Code style and conventions
- Test coverage
- Documentation quality
- Bug potential

---

## Communication Protocols

### Communication Channels

#### Slack Channels
```
#aiw3-nft-dev           # Development discussions
#aiw3-nft-alerts        # CI/CD and monitoring alerts
#aiw3-nft-releases      # Release coordination
#aiw3-nft-support       # Production issues and support
#aiw3-nft-general       # General project updates
```

#### Channel Usage Guidelines
- **#aiw3-nft-dev**: Technical discussions, architecture decisions, code questions
- **#aiw3-nft-alerts**: Automated alerts only (no discussions)
- **#aiw3-nft-releases**: Release planning, deployment coordination
- **#aiw3-nft-support**: Production issues, user reports, incident response
- **#aiw3-nft-general**: Project updates, announcements, team coordination

### Communication Standards

#### Response Time Expectations
| Priority | Response Time | Escalation |
|----------|---------------|------------|
| **Critical** (Production down) | 15 minutes | Immediate phone call |
| **High** (Feature blocked) | 2 hours | Slack mention |
| **Medium** (General question) | 4 hours | Follow-up message |
| **Low** (Documentation/planning) | 24 hours | Weekly standup |

#### Escalation Path
```
Developer → Senior Developer → Tech Lead → Engineering Manager
```

#### Status Updates
**Daily Standup Format**:
- **Yesterday**: What did you accomplish?
- **Today**: What will you work on?
- **Blockers**: What's preventing progress?
- **Help Needed**: What support do you need?

---

## Meeting & Planning Cadences

### Regular Meetings

#### Daily Standup
- **When**: Every weekday, 9:00 AM
- **Duration**: 15 minutes
- **Participants**: All developers, QA, Product Owner
- **Format**: Round-robin status updates
- **Tool**: Slack huddle or video call

#### Sprint Planning
- **When**: Every 2 weeks, Monday 2:00 PM
- **Duration**: 2 hours
- **Participants**: Development team, Product Owner, Scrum Master
- **Agenda**:
  - Review previous sprint
  - Plan upcoming sprint
  - Estimate story points
  - Identify dependencies and risks

#### Sprint Review & Retrospective
- **When**: Every 2 weeks, Friday 3:00 PM
- **Duration**: 1.5 hours
- **Participants**: Development team, stakeholders
- **Agenda**:
  - Demo completed features
  - Review sprint metrics
  - Retrospective discussion
  - Action items for improvement

#### Architecture Review
- **When**: Weekly, Wednesday 10:00 AM
- **Duration**: 1 hour
- **Participants**: Senior developers, Tech Lead, Architects
- **Agenda**:
  - Review architectural decisions
  - Discuss technical debt
  - Plan refactoring efforts
  - Address technical blockers

### Ad-Hoc Meetings

#### Technical Deep Dive
- **Trigger**: Complex technical decisions needed
- **Duration**: 1-2 hours
- **Participants**: Relevant developers and architects
- **Outcome**: Technical decision document

#### Incident Response
- **Trigger**: Production issues
- **Duration**: Until resolved
- **Participants**: On-call engineer, relevant developers, manager
- **Outcome**: Incident report and action items

---

## Documentation Standards

### Documentation Requirements

#### When to Update Documentation
- **Before** implementing new features
- **During** architectural changes
- **After** bug fixes that affect documented behavior
- **Immediately** after API changes

#### Documentation Types & Ownership

| Document Type | Owner | Review Required | Update Frequency |
|---------------|-------|-----------------|------------------|
| **System Design** | Tech Lead | Architecture Review | Major changes |
| **API Documentation** | Backend Developer | Code Review | Every API change |
| **User Guides** | Product Owner | Team Review | Feature releases |
| **Deployment Guides** | DevOps Engineer | Security Review | Infrastructure changes |
| **Best Practices** | Senior Developer | Team Consensus | Quarterly |

### Documentation Review Process

#### 1. Documentation PR Process
```markdown
## Documentation Change Request

### Type of Change
- [ ] New documentation
- [ ] Update existing documentation
- [ ] Fix documentation errors
- [ ] Remove outdated documentation

### Impact Assessment
- [ ] Affects developer onboarding
- [ ] Changes API contracts
- [ ] Updates deployment procedures
- [ ] Modifies team processes

### Review Requirements
- [ ] Technical accuracy verified
- [ ] Cross-references updated
- [ ] Related documents updated
- [ ] Stakeholders notified
```

#### 2. Documentation Quality Standards
- **Accuracy**: All technical details verified against code
- **Completeness**: All necessary information included
- **Clarity**: Written for target audience
- **Currency**: Up-to-date with latest changes
- **Cross-references**: Links to related documents maintained

---

## Issue Tracking & Project Management

### Jira Workflow

#### Issue Types
- **Epic**: Large feature or initiative
- **Story**: User-facing functionality
- **Task**: Development work
- **Bug**: Defects and issues
- **Spike**: Research or investigation

#### Issue States
```
To Do → In Progress → Code Review → Testing → Done
```

#### Issue Naming Conventions
```
NFT-123: Implement NFT claiming API endpoint
NFT-124: Fix Solana RPC timeout handling
NFT-125: Research IPFS gateway redundancy options
```

### Sprint Management

#### Story Point Estimation
- **1 point**: Simple changes, <4 hours
- **2 points**: Small features, 4-8 hours
- **3 points**: Medium features, 1-2 days
- **5 points**: Large features, 2-3 days
- **8 points**: Complex features, 3-5 days
- **13+ points**: Epic, needs breakdown

#### Sprint Capacity Planning
- **Developer capacity**: 6-8 points per sprint
- **Senior developer capacity**: 8-10 points per sprint
- **Buffer for bugs/support**: 20% of total capacity

### Definition of Done

#### For User Stories
- [ ] Acceptance criteria met
- [ ] Code reviewed and approved
- [ ] Unit tests written and passing
- [ ] Integration tests passing
- [ ] Documentation updated
- [ ] Security review completed (if applicable)
- [ ] Performance impact assessed
- [ ] Deployed to staging environment
- [ ] QA testing completed
- [ ] Product Owner acceptance

#### For Bugs
- [ ] Root cause identified
- [ ] Fix implemented and tested
- [ ] Regression tests added
- [ ] Code reviewed
- [ ] Deployed to production
- [ ] Verification completed
- [ ] Post-mortem completed (if critical)

---

## Release & Deployment Coordination

### Release Planning

#### Release Types
- **Major Release** (v1.0.0): New features, breaking changes
- **Minor Release** (v1.1.0): New features, backward compatible
- **Patch Release** (v1.1.1): Bug fixes only
- **Hotfix Release** (v1.1.2): Critical production fixes

#### Release Schedule
- **Major releases**: Quarterly
- **Minor releases**: Monthly
- **Patch releases**: As needed
- **Hotfix releases**: Emergency only

### Deployment Process

#### Pre-Deployment Checklist
- [ ] All tests passing
- [ ] Security scan completed
- [ ] Performance benchmarks met
- [ ] Database migrations tested
- [ ] Rollback plan prepared
- [ ] Monitoring alerts configured
- [ ] Team notified of deployment window

#### Deployment Coordination
1. **Deployment Lead** announces deployment window
2. **Code freeze** 24 hours before deployment
3. **Pre-deployment meeting** to review checklist
4. **Deployment execution** with team monitoring
5. **Post-deployment verification** and monitoring
6. **Go/No-Go decision** within 30 minutes
7. **Rollback execution** if issues detected

#### Post-Deployment
- [ ] Health checks verified
- [ ] Key metrics monitored
- [ ] User feedback collected
- [ ] Issues documented and prioritized
- [ ] Lessons learned captured

---

## Knowledge Sharing Practices

### Knowledge Sharing Sessions

#### Tech Talks
- **Frequency**: Bi-weekly, Friday 4:00 PM
- **Duration**: 30 minutes
- **Format**: Presentation + Q&A
- **Topics**: New technologies, architecture decisions, lessons learned

#### Code Walkthroughs
- **Frequency**: Weekly, Thursday 11:00 AM
- **Duration**: 45 minutes
- **Format**: Live code review
- **Focus**: Complex features, design patterns, best practices

#### Architecture Decision Records (ADRs)
```markdown
# ADR-001: Use Standard Solana Programs Only

## Status
Accepted

## Context
Need to decide on smart contract strategy for NFT system.

## Decision
Use only standard Solana programs (SPL Token, Metaplex) without custom contracts.

## Consequences
- Reduced development complexity
- Better ecosystem compatibility
- Lower security risk
- Limited customization options
```

### Documentation Practices

#### Code Documentation
- **Inline comments**: Explain complex logic
- **Function documentation**: JSDoc format
- **API documentation**: OpenAPI/Swagger
- **Architecture documentation**: System design docs

#### Knowledge Base
- **Confluence/Wiki**: Team knowledge base
- **Runbooks**: Operational procedures
- **Troubleshooting guides**: Common issues and solutions
- **Onboarding guides**: New team member resources

---

## Quality Assurance Coordination

### QA Process Integration

#### Development-QA Handoff
1. **Feature completion** notification to QA
2. **Test environment** deployment
3. **Test data setup** by development team
4. **QA testing** execution
5. **Bug reporting** and triage
6. **Fix verification** and sign-off

#### Testing Coordination
- **Test planning**: QA involved in sprint planning
- **Test case review**: Development team reviews test cases
- **Automation**: Shared responsibility for test automation
- **Environment management**: Coordinated test environment updates

### Bug Triage Process

#### Bug Priority Levels
- **P0 - Critical**: Production down, data loss
- **P1 - High**: Major feature broken, security issue
- **P2 - Medium**: Minor feature issue, performance problem
- **P3 - Low**: Cosmetic issue, enhancement request

#### Triage Meeting
- **Frequency**: Daily during active development
- **Participants**: Tech Lead, QA Lead, Product Owner
- **Duration**: 15 minutes
- **Outcome**: Bug priority and assignment decisions

---

## Emergency Response Procedures

### Incident Response Team
- **Incident Commander**: On-call engineer
- **Technical Lead**: Senior developer familiar with affected system
- **Communication Lead**: Project manager or team lead
- **Subject Matter Expert**: Developer who worked on affected feature

### Response Process

#### 1. Incident Detection
- **Monitoring alerts**: Automated detection
- **User reports**: Support channel notifications
- **Team discovery**: Developer identifies issue

#### 2. Initial Response (Within 15 minutes)
- [ ] Acknowledge incident in #aiw3-nft-support
- [ ] Assess severity and impact
- [ ] Activate incident response team
- [ ] Create incident tracking ticket
- [ ] Begin investigation

#### 3. Communication Protocol
- **Internal updates**: Every 30 minutes in Slack
- **Stakeholder updates**: Every hour via email
- **User communication**: Status page updates
- **Escalation**: Manager notification for P0/P1 issues

#### 4. Resolution and Follow-up
- [ ] Implement fix or workaround
- [ ] Verify resolution
- [ ] Monitor for recurrence
- [ ] Update stakeholders
- [ ] Schedule post-mortem
- [ ] Document lessons learned

### Post-Mortem Process

#### Post-Mortem Meeting
- **Timing**: Within 48 hours of resolution
- **Duration**: 1 hour
- **Participants**: Incident response team + stakeholders
- **Outcome**: Action items for prevention

#### Post-Mortem Report Template
```markdown
# Incident Post-Mortem: [Date] - [Brief Description]

## Summary
Brief description of the incident and impact.

## Timeline
Chronological sequence of events.

## Root Cause Analysis
What caused the incident and why it wasn't caught earlier.

## Impact Assessment
User impact, business impact, and system impact.

## Action Items
- [ ] Immediate fixes
- [ ] Process improvements
- [ ] Monitoring enhancements
- [ ] Documentation updates

## Lessons Learned
What we learned and how to prevent similar incidents.
```

---

## Related Documents

### Technical Documentation
- [AIW3 NFT Best Practices](./AIW3-NFT-Best-Practices.md) - Technical implementation standards
- [AIW3 NFT System Design](./AIW3-NFT-System-Design.md) - Architecture overview
- [AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md) - Development guidelines

### Process Documentation
- [AIW3 NFT Testing Strategy](./AIW3-NFT-Testing-Strategy.md) - Testing approach and standards
- [AIW3 NFT Deployment Guide](./AIW3-NFT-Deployment-Guide.md) - Deployment procedures
- [AIW3 NFT Error Handling Reference](./AIW3-NFT-Error-Handling-Reference.md) - Error management
